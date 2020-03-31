/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/acme/autocert"

	"github.com/app-nerds/corona-ramp-up/api/cache"
	"github.com/app-nerds/corona-ramp-up/api/collate"
	"github.com/app-nerds/corona-ramp-up/api/graph"
	"github.com/app-nerds/corona-ramp-up/api/jhucsse"
	"github.com/app-nerds/corona-ramp-up/api/version"
	"github.com/app-nerds/corona-ramp-up/assets"
)

/*
IApplication defines an interface for running the main application
*/
type IApplication interface {
	Start() chan os.Signal
	Stop()
}

/*
Application provides an implementation of the main application
*/
type Application struct {
	Config          *viper.Viper
	Cron            *cron.Cron
	CronEntryID     cron.EntryID
	HTTPServer      *echo.Echo
	Logger          *logrus.Entry
	ReportCache     cache.ICache
	ShutdownContext context.Context

	/*
	 * Controllers
	 */
	GraphAssemblerController graph.IGraphAssemblerController
	VersionController        *version.VersionController

	/*
	 * Services
	 */
	ReportRetriever     jhucsse.ReportRetriever
	LineSeriesAssembler graph.GraphDataAssembler
}

/*
NewApplication is the factory method to create a new Application
*/
func NewApplication(shutdownContext context.Context, logger *logrus.Entry, config *viper.Viper) *Application {
	result := &Application{
		Config:          config,
		Logger:          logger,
		ReportCache:     cache.NewCache(),
		ShutdownContext: shutdownContext,
	}

	result.setupServices()
	result.setupCache()
	result.setupControllers()
	result.setupHandlers()

	return result
}

/*
handleMainPage is what serves the primary HTML container. Modify this to include
new scripts, CSS, or change the title
*/
func (a *Application) handleMainPage(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<meta http-equiv="X-UA-Compatible" content="IE=edge" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />

	<title>COVID-19 Ramp Up Comparison</title>

	<link rel="stylesheet" type="text/css" href="/app/assets/bootstrap/css/bootstrap.min.css" />
	<link rel="stylesheet" type="text/css" href="/app/assets/corona/css/styles.css" />
	<link rel="stylesheet" type="text/css" href="/app/assets/vue-loading-overlay/vue-loading-overlay.css" />
	<link rel="stylesheet" type="text/css" href="/app/assets/syncfusion/bootstrap4.css" />
</head>

<body>
	<div id="app"></div>

	<script src="/app/assets/babel/babel.min.js"></script>
	<script src="/app/assets/moment/moment.min.js"></script>
	<script src="/app/assets/vue/vue-2.6.10.js"></script>
	<script src="/app/assets/vue-router/vue-router-3.1.3.min.js"></script>
	<script src="/app/assets/vue-resource/vue-resource-1.5.1.min.js"></script>
	<script src="/app/assets/jquery/jquery-3.4.1.min.js"></script>
	<script src="/app/assets/popper/popper.min.js"></script>
	<script src="/app/assets/bootstrap/js/bootstrap.min.js"></script>
	<script src="/app/assets/vue-loading-overlay/vue-loading-overlay.js"></script>
	<script src="/app/assets/syncfusion/ej2-vue.min.js"></script>

	<script src="/app/main.js" type="module"></script>
	</body>
</html>
`)
}

func (a *Application) setupCache() {
	var err error
	var report jhucsse.ReportRecordCollection

	reportGetter := func() {
		a.Logger.Info("Retrieving COVID-19 report")
		if report, err = a.ReportRetriever.Retrieve(); err != nil {
			a.Logger.WithError(err).Fatal("Initial cache load failed!")
		}

		a.Logger.Debug("Storing report into cache")
		a.ReportCache.StoreReport(report)
	}

	reportGetter()

	a.Cron = cron.New()
	a.CronEntryID, _ = a.Cron.AddFunc(a.Config.GetString("dailyReport.schedule"), reportGetter)
	a.Cron.Start()
}

/*
setupControllers is where you will initialize controllers that handle
your API routes
*/
func (a *Application) setupControllers() {
	a.GraphAssemblerController = graph.NewGraphAssemblerController(graph.GraphAssemblerControllerConfig{
		Config:              a.Config,
		LineSeriesAssembler: a.LineSeriesAssembler,
		Logger:              a.Logger.WithField("who", "GraphAssemblerController"),
	})

	a.VersionController = version.NewVersionController(&version.VersionControllerConfig{
		Config: a.Config,
	})
}

/*
setupHandlers is where API routes are managed. Here you wire up a URL
route to a controller function
*/
func (a *Application) setupHandlers() {
	a.HTTPServer = echo.New()
	a.HTTPServer.HideBanner = true
	a.HTTPServer.Use(middleware.CORS())

	api := a.HTTPServer.Group("/api")

	a.HTTPServer.GET("/app/*", echo.WrapHandler(http.FileServer(assets.Assets)))
	a.HTTPServer.GET("/", a.handleMainPage)

	api.GET("/regions", a.GraphAssemblerController.GetRegions)
	api.GET("/startpoints", a.GraphAssemblerController.GetStartPoints)
	api.POST("/lineseries", a.GraphAssemblerController.GetLineSeriesData)
	api.GET("/version", a.VersionController.GetVersion)
}

/*
setupServices is where you will initialize an services that are
dependencies to controllers
*/
func (a *Application) setupServices() {
	// a.ReportRetriever = jhucsse.NewLocalReportRetriever(jhucsse.ReportRetrieverConfig{
	// 	Config: a.Config,
	// 	Logger: a.Logger,
	// })

	a.ReportRetriever = jhucsse.NewGithubReportRetriever(jhucsse.ReportRetrieverConfig{
		Config: a.Config,
		Logger: a.Logger,
	})

	a.LineSeriesAssembler = graph.NewLineSeriesAssembler(graph.GraphDataAssemblerConfig{
		Collator: collate.NewCollator(collate.CollatorConfig{
			ReportCache: a.ReportCache,
		}),
	})
}

/*
Start begins serving the API application
*/
func (a *Application) Start() chan os.Signal {
	if a.Config.GetBool("server.autoSSL") {
		a.HTTPServer.AutoTLSManager.Cache = autocert.DirCache("./.certs")
		a.HTTPServer.AutoTLSManager.HostPolicy = autocert.HostWhitelist(a.Config.GetStringSlice("server.autoSSLWhitelist")...)
		a.HTTPServer.AutoTLSManager.Prompt = autocert.AcceptTOS
		a.HTTPServer.AutoTLSManager.Email = a.Config.GetString("server.autoSSLEmail")

		ipSplit := strings.Split(a.Config.GetString("server.host"), ":")
		go http.ListenAndServe(fmt.Sprintf("%s:80", ipSplit[0]), a.HTTPServer.AutoTLSManager.HTTPHandler(nil))
	}

	go func() {
		var err error

		a.Logger.WithFields(logrus.Fields{
			"host":          a.Config.GetString("server.host"),
			"serverVersion": a.Config.GetString("server.version"),
			"debug":         a.Config.GetBool("server.debug"),
			"logLevel":      a.Config.GetString("fireplace.loglevel"),
		}).Infof("Starting")

		if a.Config.GetBool("server.autoSSL") {
			err = a.HTTPServer.StartAutoTLS(a.Config.GetString("server.host"))
		} else {
			err = a.HTTPServer.Start(a.Config.GetString("server.host"))
		}

		if err != http.ErrServerClosed {
			a.Logger.WithError(err).Fatalf("Unable to start application")
		} else {
			a.Logger.Infof("Shutting down the server...")
		}
	}()

	/*
	 * Setup shutdown handler
	 */
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	return quit
}

/*
Stop halt API server execution. It waits for 10 seconds, and if
the server has not stopped a panic is thrown
*/
func (a *Application) Stop() {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	a.Cron.Stop()

	if err = a.HTTPServer.Shutdown(ctx); err != nil {
		a.Logger.WithError(err).Errorf("There was an error shutting down the server")
	}
}
