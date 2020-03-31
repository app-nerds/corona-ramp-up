/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package jhucsse

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ReportRetrieverConfig struct {
	Config *viper.Viper
	Logger *logrus.Entry
}
