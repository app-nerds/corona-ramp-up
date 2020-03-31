/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package jhucsse

import (
	"fmt"
)

var ErrNoReportData = fmt.Errorf("no report data")
var ErrReadingReportDataCSV = fmt.Errorf("error reading CSV report data")
var ErrInvalidStartDate = fmt.Errorf("invalid start date for getting daily report data")
