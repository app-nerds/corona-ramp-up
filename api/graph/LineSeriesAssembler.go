/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package graph

import (
	"github.com/app-nerds/corona-ramp-up/api/collate"
	"github.com/app-nerds/corona-ramp-up/api/jhucsse"
)

type LineSeriesAssembler struct {
	Collator collate.ICollator
}

func NewLineSeriesAssembler(config GraphDataAssemblerConfig) *LineSeriesAssembler {
	return &LineSeriesAssembler{
		Collator: config.Collator,
	}
}

func (a *LineSeriesAssembler) Assemble(regions []collate.Region, startPoint collate.StartPoint) GraphDataCollection {
	result := make(GraphDataCollection, 0, 10)

	for _, region := range regions {
		reports := a.Collator.Collate(region, startPoint)
		effectiveStartDate := a.Collator.GetEffectiveStartDate(region, startPoint)
		shutdownDescription := a.Collator.GetShutdownDescription(region, startPoint)
		seriesDataCollection := a.assembleSeriesData(reports)

		graphData := &GraphData{
			Region:              string(region),
			EffectiveStartDate:  effectiveStartDate,
			SeriesData:          seriesDataCollection,
			ShutdownDescription: shutdownDescription,
		}

		result = append(result, graphData)
	}

	return result
}

func (a *LineSeriesAssembler) assembleSeriesData(reports jhucsse.ReportRecordCollection) SeriesDataCollection {
	result := make(SeriesDataCollection, 0, 500)

	for index, report := range reports {
		seriesData := &SeriesData{
			DayNumber:      index,
			ConfirmedCases: report.Confirmed,
		}

		result = append(result, seriesData)
	}

	return result
}
