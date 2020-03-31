/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package collate

import (
	"time"

	"github.com/app-nerds/corona-ramp-up/api/cache"
	"github.com/app-nerds/corona-ramp-up/api/jhucsse"
)

type ICollator interface {
	Collate(region Region, startPoint StartPoint) jhucsse.ReportRecordCollection
	GetEffectiveStartDate(region Region, startPoint StartPoint) time.Time
	GetShutdownDescription(region Region, startPoint StartPoint) string
}

type Collator struct {
	ReportCache cache.ICache
}

type CollatorConfig struct {
	ReportCache cache.ICache
}

func NewCollator(config CollatorConfig) *Collator {
	return &Collator{
		ReportCache: config.ReportCache,
	}
}

func (c *Collator) Collate(region Region, startPoint StartPoint) jhucsse.ReportRecordCollection {
	reports := c.ReportCache.GetReport()
	return SliceRegion(reports, region, startPoint)
}

func (c *Collator) GetEffectiveStartDate(region Region, startPoint StartPoint) time.Time {
	return CollatorData[region][startPoint].DateTime
}

func (c *Collator) GetShutdownDescription(region Region, startPoint StartPoint) string {
	return CollatorData[region][startPoint].ShutdownDescription
}
