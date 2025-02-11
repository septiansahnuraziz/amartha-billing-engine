package utils

import (
	"fmt"
	"time"
)

// ParseDurationWithDefault parses a duration string and returns the parsed duration.
// If the parsing fails, it returns the default duration provided.
func ParseDurationWithDefault(input string, defaultDuration time.Duration) time.Duration {
	parsedDuration, err := time.ParseDuration(input)
	if err != nil {
		return defaultDuration
	}

	return parsedDuration
}

// ParseDate parses a date string using the specified layout and returns the corresponding time.Time value.
// If the parsing fails, it returns the zero time value.
func ParseDate(layout, value string) time.Time {
	date, _ := time.Parse(layout, value)

	return date
}

// ParseDatetimeToRFC3339 formats the provided time value to the RFC3339 format.
func ParseDatetimeToRFC3339(inputTime *time.Time) string {
	return inputTime.Format(time.RFC3339)
}

func GetNowTime() time.Time {
	return time.Now()
}

func GetNowTimeFromAsiaJKT() time.Time {
	jakartaLocation, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(jakartaLocation)
}

func InMinuteTimeRange(startTime time.Time, stopTime uint) bool {
	return time.Since(startTime) >= time.Duration(stopTime)*time.Minute
}

func GetNowTimeRFC3339() string {
	times := time.Now()

	return times.Format(time.RFC3339)
}

func GetDate(times time.Time) time.Time {
	date := time.Date(times.Year(), times.Month(), times.Day(), 0, 0, 0, 0, time.UTC)

	return date
}

func GetTimeDuration(param int) time.Duration {
	return time.Duration(param)
}

func AddTime(currentTime time.Time, additionalTime int, unit string) time.Time {
	var timeUnit time.Duration

	switch unit {
	case "second":
		timeUnit = time.Second
	case "minute":
		timeUnit = time.Minute
	case "hour":
		timeUnit = time.Hour
	}

	return currentTime.Add(time.Duration(additionalTime) * timeUnit)
}

func SubTime(currentTime time.Time, subtractionTime int, unit string) time.Time {
	var timeUnit time.Duration

	switch unit {
	case "second":
		timeUnit = time.Second
	case "minute":
		timeUnit = time.Minute
	case "hour":
		timeUnit = time.Hour
	}

	return currentTime.Add(time.Duration(-subtractionTime) * timeUnit)
}

func GetTomorrowDate(currentTime time.Time) time.Time {
	yyyy, mm, dd := currentTime.Date()
	tomorrow := time.Date(yyyy, mm, dd+1, 0, 0, 0, 0, currentTime.Location())
	return tomorrow
}

// BeginningOfDay returns the start of the given day
func BeginningOfDay(t time.Time) time.Time {
	jakartaLocation, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, jakartaLocation)
}

// EndOfDay returns the end of the given day
func EndOfDay(t time.Time) time.Time {
	jakartaLocation, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, jakartaLocation)
}

// BeginningOfMonth returns the start of the month for the given time
func BeginningOfMonth(t time.Time) time.Time {
	jakartaLocation, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, jakartaLocation)
}

// EndOfMonth returns the end of the month for the given time
func EndOfMonth(t time.Time) time.Time {
	// Add 1 month to the first day of the current month, then subtract 1 nanosecond
	return BeginningOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginningOfYear returns the start of the year for the given time
func BeginningOfYear(t time.Time) time.Time {
	jakartaLocation, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, jakartaLocation)
}

// EndOfYear returns the end of the year for the given time
func EndOfYear(t time.Time) time.Time {
	jakartaLocation, _ := time.LoadLocation("Asia/Jakarta")
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, jakartaLocation)
}

// Yesterday returns the start of yesterday
func Yesterday() time.Time {
	return BeginningOfDay(time.Now().AddDate(0, 0, -1))
}

// Tomorrow returns the start of tomorrow
func Tomorrow() time.Time {
	return BeginningOfDay(time.Now().AddDate(0, 0, 1))
}

// GetOneMonthPastRange returns the date range for 1 month before the given time
func GetOneMonthPastRange(now time.Time) (startAt, endAt time.Time) {
	jakartaLocation, _ := time.LoadLocation("Asia/Jakarta")
	endAt = now.In(jakartaLocation)
	startAt = now.AddDate(0, -1, 0).Add(time.Nanosecond).In(jakartaLocation)
	return startAt, endAt
}

// DateFormats contains a list of date formats that we'll try to parse
var DateFormats = []string{
	// Standard formats
	time.RFC3339,
	time.RFC822,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339Nano,

	// ISO 8601 variants
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05",
	"2006-01-02T15:04:05-07:00",
	"2006-01-02T15:04:05.999999999Z07:00",

	// Long date formats with month names
	"02 January 2006",
	"2 January 2006",
	"January 02, 2006",
	"January 2, 2006",
	"02 Jan 2006",
	"2 Jan 2006",
	"Jan 02, 2006",
	"Jan 2, 2006",

	// Long date formats with time
	"02 January 2006 15:04",
	"02 January 2006 15:04:05",
	"2 January 2006 15:04",
	"2 January 2006 15:04:05",
	"January 02, 2006 15:04",
	"January 02, 2006 15:04:05",
	"02 Jan 2006 15:04",
	"02 Jan 2006 15:04:05",

	// Common date and time formats
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02 15:04:05 -0700",
	"2006-01-02 15:04:05 MST",

	// Date only formats
	"2006-01-02",
	"2006/01/02",
	"02/01/2006",
	"02-01-2006",
	"01/02/2006",
	"01-02-2006",

	// Excel/Spreadsheet style formats
	"02/01/2006 15:04",    // DD/MM/YYYY HH:mm
	"30/10/2024 10:30",    // Specific example
	"02-01-2006 15:04",    // DD-MM-YYYY HH:mm
	"01/02/2006 15:04:05", // MM/DD/YYYY HH:mm:ss
	"01-02-2006 15:04:05", // MM-DD-YYYY HH:mm:ss

	// Additional international formats
	"2006年01月02日",         // Japanese format
	"02.01.2006",          // European format with dots
	"02.01.2006 15:04",    // European format with time
	"02.01.2006 15:04:05", // European format with seconds

	// Additional time formats
	"15:04:05",
	"15:04",
	"3:04 PM",
	"3:04:05 PM",
}

// ParseFlexibleDate attempts to parse the input string using various date formats
func ParseFlexibleDate(dateStr string) (time.Time, error) {
	for _, format := range DateFormats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}

func ParseTimeGetEmptyStringWhenZero(t *time.Time, toLayout string) string {
	if t == nil || t.IsZero() {
		return ""
	}
	return t.Format(toLayout)
}

func ParseDateInLocation(layout, value string) time.Time {
	date, _ := time.ParseInLocation(layout, value, time.Local)

	return date
}

func GetCurrentWeek(startDate time.Time) int {
	today := time.Now()
	duration := today.Sub(startDate)
	weekNumber := int(duration.Hours()/24/7) + 1
	return weekNumber
}

func GetCurrentWeekBilling(startDate time.Time, totalWeeks int) (int, bool) {
	weekNumber := GetCurrentWeek(startDate)

	if weekNumber > totalWeeks {
		return weekNumber, false
	}

	return weekNumber, true
}

func IsPastDueDate(dueDate time.Time) bool {
	today := time.Now()
	return today.After(dueDate)
}
