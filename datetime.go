package gostradamus

import (
	"fmt"
	"time"
)

// Datetime reflects time.Time and adds functionality
type DateTime time.Time

// DateTimeFromTime returns a DateTime with given time.Time
func DateTimeFromTime(time time.Time) DateTime {
	return DateTime(time)
}

// NewDateTime returns a new DateTime in the timezone given
func NewDateTime(
	year int,
	month int,
	day int,
	hour int,
	minute int,
	second int,
	nanosecond int,
	timezone Timezone,
) DateTime {
	return DateTimeFromTime(
		time.Date(
			year,
			time.Month(month),
			day,
			hour,
			minute,
			second,
			nanosecond,
			timezone.Location(),
		),
	)
}

// NewUTCDateTime returns a new DateTime with timezone in UTC
func NewUTCDateTime(
	year int,
	month int,
	day int,
	hour int,
	minute int,
	second int,
	nanosecond int,
) DateTime {
	return NewDateTime(
		year,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		UTC,
	)
}

// NewLocalDateTime creates a new DateTime in Local (system) timezone
func NewLocalDateTime(
	year int,
	month int,
	day int,
	hour int,
	minute int,
	second int,
	nanosecond int,
) DateTime {
	return NewDateTime(
		year,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		Local(),
	)
}

// FromUnixTimestamp gets the DateTime from given unix timestamp.
// The returned Datetime has the UTC timezone
func FromUnixTimestamp(timestamp int64) DateTime {
	return DateTime(time.Unix(timestamp, 0).UTC())
}

// Now returns the current local DateTime
func Now() DateTime {
	return DateTimeFromTime(time.Now())
}

// UTCNow returns the current DateTime in UTC timezone
func UTCNow() DateTime {
	return Now().InTimezone(UTC)
}

// NowInTimezone returns the current DateTime in given timezone
func NowInTimezone(timezone Timezone) DateTime {
	return Now().InTimezone(timezone)
}

// Year of current DateTime as int
func (dt DateTime) Year() int {
	return dt.Time().Year()
}

// Month of current DateTime as int
func (dt DateTime) Month() int {
	return int(dt.Time().Month())
}

// Day of current DateTime as int
func (dt DateTime) Day() int {
	return dt.Time().Day()
}

// Hour of current DateTime as int
func (dt DateTime) Hour() int {
	return dt.Time().Hour()
}

// Minute of current DateTime as int
func (dt DateTime) Minute() int {
	return dt.Time().Minute()
}

// Second of current DateTime as int
func (dt DateTime) Second() int {
	return dt.Time().Second()
}

// Nanosecond of current DateTime as int
func (dt DateTime) Nanosecond() int {
	return dt.Time().Nanosecond()
}

// Time returns the underlying time.Time of DateTime
func (dt DateTime) Time() time.Time {
	return time.Time(dt)
}

// IsoFormat the current DateTime into ISO-8601 standard
// Example: 2017-07-14T02:40:00.000000+0200
func (dt DateTime) IsoFormat() string {
	return dt.Format(Iso8601)
}

// IsoFormatTZ the current DateTime into ISO-8601 standard and current timezone info
// Example: 2017-07-14T02:40:00.000000
func (dt DateTime) IsoFormatTZ() string {
	return dt.Format(Iso8601TZ)
}

// CTimeFormat the current DateTime to ctime standard
// Example: Sat Feb 15 12:12:12 2020
func (dt DateTime) CTimeFormat() string {
	return dt.Format(CTime)
}

// String of the DateTime
// it will use the IsoFormat
// It defines the ``native'' format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
func (dt DateTime) String() string {
	return dt.IsoFormatTZ()
}

// GoString of the DateTime
// it will use the IsoFormat
// The GoString method is used to print values passed as an operand to a %#v format.
func (dt DateTime) GoString() string {
	return dt.IsoFormatTZ()
}

// Copy the current DateTime to a new DateTime
func (dt DateTime) Copy() DateTime {
	return NewDateTime(
		dt.Year(),
		dt.Month(),
		dt.Day(),
		dt.Hour(),
		dt.Minute(),
		dt.Second(),
		dt.Nanosecond(),
		dt.Timezone(),
	)
}

// Format the current DateTime with given format to a string
func (dt DateTime) Format(format string) string {
	return formatFromTime(dt.Time(), format)
}

// Parse a string value with given format into a new DateTime
func Parse(value string, format string) (DateTime, error) {
	parsedTime, err := parseToTime(value, format, UTC)
	return DateTimeFromTime(parsedTime), err
}

// ParseInTimezone a string value with given format into a new DateTime in given timezone
func ParseInTimezone(value string, format string, timezone Timezone) (DateTime, error) {
	parsedTime, err := parseToTime(value, format, timezone)
	return DateTimeFromTime(parsedTime), err
}

// InTimezone sets the current DateTime in the given Timezone and returns a new DateTime
func (dt DateTime) InTimezone(timezone Timezone) DateTime {
	return DateTimeFromTime(dt.Time().In(timezone.Location()))
}

// IsoCalendar returns three int values with (year, month, day)
func (dt DateTime) IsoCalendar() (int, int, int) {
	return dt.Year(), dt.Month(), dt.Day()
}

// Timezone returns the Timezone of current DateTime object
func (dt DateTime) Timezone() Timezone {
	return Timezone(dt.Time().Location().String())
}

// UnixTimestamp returns the unix timestamp as int64
func (dt DateTime) UnixTimestamp() int64 {
	return dt.Time().Unix()
}

// ShiftYears adds or subtracts years in current DateTime and returns a new DateTime
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) ShiftYears(years int) DateTime {
	return DateTimeFromTime(dt.Time().AddDate(years, 0, 0))
}

// ShiftMonths adds or subtracts months
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) ShiftMonths(months int) DateTime {
	return DateTimeFromTime(dt.Time().AddDate(0, months, 0))
}

// ShiftDays adds or subtracts days
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) ShiftDays(days int) DateTime {
	return DateTimeFromTime(dt.Time().AddDate(0, 0, days))
}

// ShiftHours adds or subtracts hours
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) ShiftHours(hours int) DateTime {
	duration, _ := time.ParseDuration(fmt.Sprintf("%dh", hours))
	return DateTime(dt.Time().Add(duration))
}

// ShiftMinutes adds or subtracts minutes
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) ShiftMinutes(minutes int) DateTime {
	duration, _ := time.ParseDuration(fmt.Sprintf("%dm", minutes))
	return DateTime(dt.Time().Add(duration))
}

// ShiftSeconds adds or subtracts seconds
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) ShiftSeconds(second int) DateTime {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", second))
	return DateTime(dt.Time().Add(duration))
}

// ShiftMilliSeconds adds or subtracts milliseconds
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) ShiftMilliSeconds(millisecond int) DateTime {
	duration, _ := time.ParseDuration(fmt.Sprintf("%dms", millisecond))
	return DateTime(dt.Time().Add(duration))
}

// ShiftMicroSeconds adds or subtracts microseconds
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) ShiftMicroSeconds(microsecond int) DateTime {
	duration, _ := time.ParseDuration(fmt.Sprintf("%dus", microsecond))
	return DateTime(dt.Time().Add(duration))
}

// ShiftNanoseconds adds or subtracts nanoseconds in current DateTime and returns a new DateTime
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) ShiftNanoseconds(nanosecond int) DateTime {
	duration, _ := time.ParseDuration(fmt.Sprintf("%dns", nanosecond))
	return DateTime(dt.Time().Add(duration))
}

// Shift adds or subtracts years, months, days, hours, minutes, seconds, and nanoseconds of current DateTime and returns a new DateTime
// Add is a positive integer
// Subtract is a negative integer
func (dt DateTime) Shift(years int, months int, days int, hours int, minutes int, seconds int, nanoseconds int) DateTime {
	return dt.ShiftYears(years).
		ShiftMonths(months).
		ShiftDays(days).
		ShiftHours(hours).
		ShiftMinutes(minutes).
		ShiftSeconds(seconds).
		ShiftNanoseconds(nanoseconds)
}

// FloorYear returns a DateTime with all values to "floor" except year
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-01-01 00:00:00.00000000
func (dt DateTime) FloorYear() DateTime {
	return dt.Replace(dt.Year(), 1, 1, 0, 0, 0, 0)
}

// FloorMonth returns a DateTime with all values to "floor" except year, month
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-01 00:00:00.00000000
func (dt DateTime) FloorMonth() DateTime {
	return dt.Replace(dt.Year(), dt.Month(), 1, 0, 0, 0, 0)
}

// FloorDay returns a DateTime with all values to "floor" except year, month, day
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-12 00:00:00.00000000
func (dt DateTime) FloorDay() DateTime {
	return dt.Replace(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0)
}

// FloorHour returns a DateTime with all values to "floor" except year, month, day, hour
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-12 12:00:00.00000000
func (dt DateTime) FloorHour() DateTime {
	return dt.Replace(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), 0, 0, 0)
}

// FloorMinute returns a DateTime with all values to "floor" except year, month, day, hour, minute
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-12 12:12:00.00000000
func (dt DateTime) FloorMinute() DateTime {
	return dt.Replace(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), 0, 0)
}

// FloorSecond returns a DateTime with all values to "floor" except year, month, day, hour, minute, second
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-31 12:12:12.00000000
func (dt DateTime) FloorSecond() DateTime {
	return dt.Replace(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), 0)
}

// CeilYear returns a DateTime with all values to "ceil" except year, month, day, hour, minute, second
// Example: 2012-05-12 12:12:12.123456789 becomes 2012-12-31 23:59:59.999999999
func (dt DateTime) CeilYear() DateTime {
	return dt.Replace(dt.Year(), 12, 31, 23, 59, 59, 999999999)
}

// CeilMonth returns a DateTime with all values to "ceil" except year, month, day, hour, minute, second
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-12 23:59:59.999999999
func (dt DateTime) CeilMonth() DateTime {
	tempDateTime := dt.Copy()
	tempDateTime = tempDateTime.ShiftMonths(1)
	tempDateTime = tempDateTime.ReplaceDay(1)
	tempDateTime = tempDateTime.ShiftDays(-1)
	return dt.Replace(dt.Year(), dt.Month(), tempDateTime.Day(), 23, 59, 59, 999999999)
}

// CeilDay returns a DateTime with all values to "ceil" except year, month, day, minute, second
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-12 23:59:59.999999999
func (dt DateTime) CeilDay() DateTime {
	return dt.Replace(dt.Year(), dt.Month(), dt.Day(), 23, 59, 59, 999999999)
}

// CeilHour returns a DateTime with all values to "ceil" except year, month, day, hour
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-12 12:59:59.999999999
func (dt DateTime) CeilHour() DateTime {
	return dt.Replace(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), 59, 59, 999999999)
}

// CeilMinute returns a DateTime with all values to "ceil" except year, month, day, hour, minute
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-12 12:12:59.999999999
func (dt DateTime) CeilMinute() DateTime {
	return dt.Replace(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), 59, 999999999)
}

// CeilSecond returns a DateTime with all values to "ceil" except year, month, day, hour, minute, second
// Example: 2012-12-12 12:12:12.123456789 becomes 2012-12-12 12:12:12.999999999
func (dt DateTime) CeilSecond() DateTime {
	return dt.Replace(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), 999999999)
}

// SpanYear returns the start and end DateTime of current year span
// Example 2012-05-12 12:12:12:123456789 becomes (2012-01-01 00:00:00.000000000, 2012-12-31 23:59:59.999999999)
func (dt DateTime) SpanYear() (DateTime, DateTime) {
	return dt.FloorYear(), dt.CeilYear()
}

// SpanMonth returns the start and end DateTime of current month span
// Example 2012-12-12 12:12:12:123456789 becomes (2012-12-01 00:00:00.000000000, 2012-12-31 23:59:59.999999999)
func (dt DateTime) SpanMonth() (DateTime, DateTime) {
	return dt.FloorMonth(), dt.CeilMonth()
}

// SpanDay returns the start and end DateTime of current day span
// Example 2012-12-12 12:12:12:123456789 becomes (2012-12-12 00:00:00.000000000, 2012-12-12 23:59:59.999999999)
func (dt DateTime) SpanDay() (DateTime, DateTime) {
	return dt.FloorDay(), dt.CeilDay()
}

// SpanHour returns the start and end DateTime of current hour span
// Example 2012-12-12 12:12:12:123456789 becomes (2012-12-12 12:00:00.000000000, 2012-12-12 12:59:59.999999999)
func (dt DateTime) SpanHour() (DateTime, DateTime) {
	return dt.FloorHour(), dt.CeilHour()
}

// SpanMinute returns the start and end DateTime of current minute span
// Example 2012-12-12 12:12:12:123456789 becomes (2012-12-12 12:12:00.000000000, 2012-12-12 12:12:59.999999999)
func (dt DateTime) SpanMinute() (DateTime, DateTime) {
	return dt.FloorMinute(), dt.CeilMinute()
}

// SpanSecond returns the start and end DateTime of current second span
// Example 2012-12-12 12:12:12:123456789 becomes (2012-12-12 12:12:12.000000000, 2012-12-12 12:12:12.999999999)
func (dt DateTime) SpanSecond() (DateTime, DateTime) {
	return dt.FloorSecond(), dt.CeilSecond()
}

// WeekDay returns time.Weekday
func (dt DateTime) WeekDay() time.Weekday {
	return dt.Time().Weekday()
}

// ReplaceYear will set the year of current DateTime and returns a new DateTime
func (dt DateTime) ReplaceYear(year int) DateTime {
	return NewDateTime(
		year,
		dt.Month(),
		dt.Day(),
		dt.Hour(),
		dt.Minute(),
		dt.Second(),
		dt.Nanosecond(),
		dt.Timezone(),
	)
}

// ReplaceMonth will set the month of current DateTime and returns a new DateTime
func (dt DateTime) ReplaceMonth(month int) DateTime {
	return NewDateTime(
		dt.Year(),
		month,
		dt.Day(),
		dt.Hour(),
		dt.Minute(),
		dt.Second(),
		dt.Nanosecond(),
		dt.Timezone(),
	)
}

// ReplaceDay will set the day of current DateTime and returns a new DateTime
func (dt DateTime) ReplaceDay(day int) DateTime {
	return NewDateTime(
		dt.Year(),
		dt.Month(),
		day,
		dt.Hour(),
		dt.Minute(),
		dt.Second(),
		dt.Nanosecond(),
		dt.Timezone(),
	)
}

// ReplaceHour will set the hour of current DateTime and returns a new DateTime
func (dt DateTime) ReplaceHour(hour int) DateTime {
	return NewDateTime(
		dt.Year(),
		dt.Month(),
		dt.Day(),
		hour,
		dt.Minute(),
		dt.Second(),
		dt.Nanosecond(),
		dt.Timezone(),
	)
}

// ReplaceMinute will set the minute of current DateTime and returns a new DateTime
func (dt DateTime) ReplaceMinute(minute int) DateTime {
	return NewDateTime(
		dt.Year(),
		dt.Month(),
		dt.Day(),
		dt.Hour(),
		minute,
		dt.Second(),
		dt.Nanosecond(),
		dt.Timezone(),
	)
}

// ReplaceSecond will set the second of current DateTime and returns a new DateTime
func (dt DateTime) ReplaceSecond(second int) DateTime {
	return NewDateTime(
		dt.Year(),
		dt.Month(),
		dt.Day(),
		dt.Hour(),
		dt.Minute(),
		second,
		dt.Nanosecond(),
		dt.Timezone(),
	)
}

// ReplaceNanosecond will set the nanosecond of current DateTime and returns a new DateTime
func (dt DateTime) ReplaceNanosecond(nanosecond int) DateTime {
	return NewDateTime(
		dt.Year(),
		dt.Month(),
		dt.Day(),
		dt.Hour(),
		dt.Minute(),
		dt.Second(),
		nanosecond,
		dt.Timezone(),
	)
}

// Replace will set the year, month, day, hour, minute, second, and nanosecond of current DateTime and returns a new DateTime
func (dt DateTime) Replace(year int, month int, day int, hour int, minute int, second int, nanosecond int) DateTime {
	return dt.ReplaceYear(year).
		ReplaceMonth(month).
		ReplaceDay(day).
		ReplaceHour(hour).
		ReplaceMinute(minute).
		ReplaceSecond(second).
		ReplaceNanosecond(nanosecond)
}

// IsBetween checks if current DateTime is between start and end DateTimes
func (dt DateTime) IsBetween(start DateTime, end DateTime) bool {
	return dt.Time().After(start.Time()) && dt.Time().Before(end.Time())
}
