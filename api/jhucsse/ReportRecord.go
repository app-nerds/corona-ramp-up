/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package jhucsse

import (
	"time"
)

/*
ReportRecord represents a single record from a John Hopkins
University Center for Systems Science and Engineering COVID-19
tracking data
*/
type ReportRecord struct {
	Region     string    `json:"region"`
	LastUpdate time.Time `json:"lastUpdate"`
	Confirmed  int64     `json:"confirmed"`
	Deaths     int64     `json:"deaths"`
	Recovered  int64     `json:"recovered"`
}

/*
ReportRecordCollection is a slice of records from a single day's
reporting CSV file
*/
type ReportRecordCollection []*ReportRecord

