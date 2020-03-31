/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewConfig(serverVersion string) *viper.Viper {
	var err error

	result := viper.New()
	result.Set("server.version", serverVersion)

	result.SetDefault("server.host", "0.0.0.0:8080")
	result.SetDefault("server.loglevel", "debug")
	result.SetDefault("server.autoSSL", false)
	result.SetDefault("server.autoSSLWhitelist", []string{})
	result.SetDefault("server.autoSSLEmail", "")

	result.SetDefault("dailyReport.baseURL", "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports")
	result.SetDefault("dailyReport.schedule", "0 1 * * *")

	result.SetConfigName("config")
	result.SetConfigType("yaml")
	result.AddConfigPath("/opt/corona-ramp-up")
	result.AddConfigPath("C:\\corona-ramp-up")
	result.AddConfigPath("$HOME/.corona-ramp-up")
	result.AddConfigPath(".")

	if err = result.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Error reading configuration file")
			panic(err)
		}
	}

	return result
}
