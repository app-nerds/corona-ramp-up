/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package graph

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/app-nerds/corona-ramp-up/api/collate"
)

type IGraphAssemblerController interface {
	GetRegions(ctx echo.Context) error
	GetStartPoints(ctx echo.Context) error
	GetLineSeriesData(ctx echo.Context) error
}

type GraphAssemblerController struct {
	Config              *viper.Viper
	LineSeriesAssembler GraphDataAssembler
	Logger              *logrus.Entry
}

type GraphAssemblerControllerConfig struct {
	Config              *viper.Viper
	LineSeriesAssembler GraphDataAssembler
	Logger              *logrus.Entry
}

type GraphDataRequest struct {
	Regions    []collate.Region   `json:"regions"`
	StartPoint collate.StartPoint `json:"startPoint"`
}

func NewGraphAssemblerController(config GraphAssemblerControllerConfig) *GraphAssemblerController {
	return &GraphAssemblerController{
		Config:              config.Config,
		LineSeriesAssembler: config.LineSeriesAssembler,
		Logger:              config.Logger,
	}
}

func (c *GraphAssemblerController) GetRegions(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, collate.CollatorRegions)
}

func (c *GraphAssemblerController) GetStartPoints(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, collate.CollatorStartPoints)
}

func (c *GraphAssemblerController) GetLineSeriesData(ctx echo.Context) error {
	var err error
	var graphDataCollection GraphDataCollection

	request := &GraphDataRequest{}

	if err = ctx.Bind(&request); err != nil {
		c.Logger.WithError(err).Error("Error binding request in GetLineSeriesData")
		return ctx.String(http.StatusBadRequest, "Invalid request")
	}

	graphDataCollection = c.LineSeriesAssembler.Assemble(request.Regions, request.StartPoint)
	return ctx.JSON(http.StatusOK, graphDataCollection)
}
