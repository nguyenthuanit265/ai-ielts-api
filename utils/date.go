package utils

import (
	"fmt"
	"time"
)

func ConvertToTime(i interface{}) (time.Time, error) {
	if s, ok := i.(string); ok {
		t, err := time.Parse(DATE_TIME_FORMAT_YYYY_MM_DD_HH_mm_SS, s)
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}
	return time.Time{}, fmt.Errorf("unable to convert to time.Time")
}

func IsZeroTime(t time.Time) bool {
	zeroTime1 := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	zeroTime2 := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	if t.Equal(zeroTime1) || t.Equal(zeroTime2) {
		return true
	}
	return false
}

func StartDate(dateStr string) string {
	date, _ := time.Parse(DATE_TIME_FORMAT_YYYY_MM_DD, dateStr)
	startDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return startDate.Format(DATE_TIME_FORMAT_YYYY_MM_DD_HH_mm_SS)
}

func EndDate(dateStr string) string {
	date, _ := time.Parse(DATE_TIME_FORMAT_YYYY_MM_DD, dateStr)
	endDate := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
	return endDate.Format(DATE_TIME_FORMAT_YYYY_MM_DD_HH_mm_SS)
}

func ToTimeByTimezone(date string, timezone string) time.Time {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		ShowErrorLogs(fmt.Errorf("unable to get time zone location"))
	}
	t, _ := time.ParseInLocation(DATE_TIME_FORMAT_YYYY_MM_DD, date, loc)
	return t
}

func IsFormatDate(date string, format string) bool {
	_, err := time.Parse(format, date)
	if err != nil {
		return false
	}
	return true
}
