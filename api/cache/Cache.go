/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package cache

import (
	gocache "github.com/patrickmn/go-cache"

	"github.com/app-nerds/corona-ramp-up/api/jhucsse"
)

type ICache interface {
	GetReport() jhucsse.ReportRecordCollection
	StoreReport(report jhucsse.ReportRecordCollection)
}

type Cache struct {
	Cache *gocache.Cache
}

func NewCache() *Cache {
	return &Cache{
		Cache: gocache.New(gocache.NoExpiration, -1),
	}
}

func (c *Cache) GetReport() jhucsse.ReportRecordCollection {
	value, _ := c.Cache.Get(KeyReport)
	result, _ := value.(jhucsse.ReportRecordCollection)
	return result
}

func (c *Cache) StoreReport(report jhucsse.ReportRecordCollection) {
	c.Cache.Set(KeyReport, report, gocache.NoExpiration)
}
