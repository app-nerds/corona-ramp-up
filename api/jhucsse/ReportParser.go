/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package jhucsse

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

type ReportParser interface {
	Parse(reportData io.Reader) (ReportRecordCollection, error)
}

type CSVReportParser struct {
}

func NewDailyReportParser() *CSVReportParser {
	return &CSVReportParser{}
}

/*
Parse takes a reader to a John Hopkins Center for Systems Science and Engineering
CSV file and returns a ReportRecordCollection
*/
func (p *CSVReportParser) Parse(reportData io.Reader) (ReportRecordCollection, error) {
	var err error
	var record []string
	var headerMap map[string]int

	result := make(ReportRecordCollection, 0, 500)
	numRecords := 0

	reader := csv.NewReader(reportData)

	/*
	 * Read and map the header
	 */
	record, err = reader.Read()

	if err == io.EOF {
		return result, ErrNoReportData
	}

	headerMap = p.createHeaderMap(record)

	/*
	 * Read the remaining records and map them to a ReportRecordCollection
	 */
	for {
		record, err = reader.Read()

		if err == io.EOF {
			if numRecords == 0 {
				return result, ErrNoReportData
			}

			break
		}

		if err != nil {
			return result, ErrReadingReportDataCSV
		}

		converted := p.convertRecordToDailyReportRecord(record, headerMap)
		result = append(result, converted)
		numRecords++
	}

	return result, nil
}

func (p *CSVReportParser) createHeaderMap(record []string) map[string]int {
	headerMap := make(map[string]int)

	for index, headerValue := range record {
		headerMap[headerValue] = index
	}

	return headerMap
}

func (p *CSVReportParser) convertRecordToDailyReportRecord(record []string, headerMap map[string]int) *ReportRecord {
	var err error
	var lastUpdate time.Time
	var confirmed int64
	var deaths int64
	var recovered int64

	result := &ReportRecord{
		Region:   record[headerMap[HeaderRegion]],
	}

	if lastUpdate, err = time.Parse("1/2/06", record[headerMap[HeaderLastUpdate]]); err == nil {
		result.LastUpdate = lastUpdate
	}

	if confirmed, err = strconv.ParseInt(record[headerMap[HeaderConfirmed]], 0, 32); err == nil {
		result.Confirmed = confirmed
	}

	if deaths, err = strconv.ParseInt(record[headerMap[HeaderDeaths]], 0, 32); err == nil {
		result.Deaths = deaths
	}

	if recovered, err = strconv.ParseInt(record[headerMap[HeaderRecovered]], 0, 32); err == nil {
		result.Recovered = recovered
	}

	return result
}
