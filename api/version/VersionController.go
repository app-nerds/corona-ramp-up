/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package version

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type IVersionController interface {
	GetVersion(ctx echo.Context) error
}

type VersionController struct {
	config *viper.Viper
}

func NewVersionController(config *VersionControllerConfig) *VersionController {
	return &VersionController{
		config: config.Config,
	}
}

func (c *VersionController) GetVersion(ctx echo.Context) error {
	return ctx.String(http.StatusOK, c.config.GetString("server.version"))
}
