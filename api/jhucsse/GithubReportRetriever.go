/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package jhucsse

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type GithubReportRetriever struct {
	Config *viper.Viper
	Logger *logrus.Entry
}

func NewGithubReportRetriever(config ReportRetrieverConfig) *GithubReportRetriever {
	return &GithubReportRetriever{
		Config: config.Config,
		Logger: config.Logger,
	}
}

func (r *GithubReportRetriever) Retrieve() (ReportRecordCollection, error) {
	var err error
	var request *http.Request
	var response *http.Response

	client := &http.Client{}
	url := r.Config.GetString("dailyReport.baseURL")

	parser := NewDailyReportParser()
	result := make(ReportRecordCollection, 0, 1000)

	if request, err = http.NewRequest("GET", url, nil); err != nil {
		return result, fmt.Errorf("error creating HTTP request: %w", err)
	}

	if response, err = client.Do(request); err != nil {
		return result, fmt.Errorf("error executing HTTP request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		var body []byte
		body, _ = ioutil.ReadAll(response.Body)

		return result, fmt.Errorf("HTTP request returned %d - %s: %w", response.StatusCode, string(body), err)
	}

	if result, err = parser.Parse(response.Body); err != nil {
		r.Logger.WithError(err).WithField("url", url).Error("Error parsing report")
		return result, fmt.Errorf("error parsing report '%s': %w", url, err)
	}

	return result, nil
}
