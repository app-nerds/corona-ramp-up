/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package jhucsse

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LocalReportRetriever struct {
	Config *viper.Viper
	Logger *logrus.Entry
}

func NewLocalReportRetriever(config ReportRetrieverConfig) *LocalReportRetriever {
	return &LocalReportRetriever{
		Config: config.Config,
		Logger: config.Logger,
	}
}

func (r *LocalReportRetriever) Retrieve() (ReportRecordCollection, error) {
	var err error
	var f *os.File

	fileName := r.Config.GetString("dailyReport.baseURL")
	parser := NewDailyReportParser()
	result := make(ReportRecordCollection, 0, 1000)

	if f, err = os.Open(fileName); err != nil {
		return result, fmt.Errorf("error opening local report file '%s': %w", fileName, err)
	}

	if result, err = parser.Parse(f); err != nil {
		r.Logger.WithError(err).WithField("fileName", fileName).Error("Error parsing report")
		return result, fmt.Errorf("error parsing report '%s': %w", fileName, err)
	}

	return result, nil
}

