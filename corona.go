/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

//go:generate go run -tags=dev assets_generate.go

package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/app-nerds/corona-ramp-up/configuration"
)

func main() {
	var loglevel logrus.Level

	config := configuration.NewConfig("0.1.0")

	logger := logrus.New().WithField("who", "Corona Ramp Up")

	loglevel, _ = logrus.ParseLevel(config.GetString("server.loglevel"))
	logger.Logger.SetLevel(loglevel)

	shutdownContext, cancelFunc := context.WithCancel(context.Background())

	application := NewApplication(shutdownContext, logger, config)
	quit := application.Start()

	<-quit
	cancelFunc()

	application.Stop()
	logger.Info("Application stopped")
}
