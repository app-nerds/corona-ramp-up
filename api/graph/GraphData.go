/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package graph

import (
	"time"
)

type GraphData struct {
	Region              string               `json:"region"`
	EffectiveStartDate  time.Time            `json:"effectiveStartDate"`
	SeriesData          SeriesDataCollection `json:"seriesData"`
	ShutdownDescription string               `json:"shutdownDescription"`
}

type GraphDataCollection []*GraphData

type SeriesData struct {
	DayNumber      int   `json:"dayNumber"`
	ConfirmedCases int64 `json:"confirmedCases"`
}

type SeriesDataCollection []*SeriesData
