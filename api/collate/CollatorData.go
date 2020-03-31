/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package collate

import (
	"strings"
	"time"

	"github.com/app-nerds/corona-ramp-up/api/jhucsse"
)

type Region string
type StartPoint string

type CollatorDataStruct struct {
	DateTime            time.Time
	ShutdownDescription string
}

var CollatorRegions = []Region{
	"China",
	"US",
	"Italy",
}

var CollatorStartPoints = []StartPoint{
	"First Major Step",
	"All Data",
}

var CollatorStartDescriptions = []string{

}
var CollatorData = map[Region]map[StartPoint]CollatorDataStruct{
	"China": {
		"First Major Step": CollatorDataStruct{
			DateTime: time.Date(2020, 1, 23, 0, 0, 0, 0, time.Now().Location()),
			ShutdownDescription: "China announced a quarantine stopping all travel in and out of Wuhan",
		},
		"All Data": CollatorDataStruct{
			DateTime: time.Date(2020, 1, 22, 0, 0, 0, 0, time.Now().Location()),
		},
	},

	"US": {
		"First Major Step": CollatorDataStruct{
			DateTime: time.Date(2020, 3, 11, 0, 0, 0, 0, time.Now().Location()),
			ShutdownDescription: "The Trump administration announces travel restrictions from Europe to the US",
		},
		"All Data": CollatorDataStruct{
			DateTime: time.Date(2020, 1, 22, 0, 0, 0, 0, time.Now().Location()),
		},
	},

	"Italy": {
		"First Major Step": CollatorDataStruct{
			DateTime: time.Date(2020, 2, 25, 0, 0, 0, 0, time.Now().Location()),
			ShutdownDescription: "Italy issues complete lockdown of 10 Northern Italian towns",
		},
		"All Data": CollatorDataStruct{
			DateTime: time.Date(2020, 1, 22, 0, 0, 0, 0, time.Now().Location()),
		},
	},
}

func SliceRegion(reports jhucsse.ReportRecordCollection, region Region, startPoint StartPoint) jhucsse.ReportRecordCollection {
	result := make(jhucsse.ReportRecordCollection, 0, 1000)

	for _, report := range reports {
		regionStartPoint := CollatorData[region][startPoint].DateTime
		regionStartPoint = regionStartPoint.Add((24 * time.Hour) * -1)

		if strings.ToLower(report.Region) == strings.ToLower(string(region)) {
			if report.LastUpdate.After(regionStartPoint) {
				result = append(result, report)
			}
		}
	}

	return result
}
