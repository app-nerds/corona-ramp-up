/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package jhucsse

type ReportRetriever interface {
	Retrieve() (ReportRecordCollection, error)
}


