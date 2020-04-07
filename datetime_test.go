package gostradamus

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	actual, err := Parse("10.02.2010 14:59:53", "DD.MM.YYYY HH:mm:ss")
	assert.NoError(t, err)
	assert.Equal(t, actual, NewUTCDateTime(2010, 2, 10, 14, 59, 53, 0))

	actual, err = Parse("2017-07-14T02:40:00.000000", Iso8601)
	assert.NoError(t, err)
	assert.Equal(t, NewDateTime(2017, 7, 14, 2, 40, 0, 0, UTC), actual)
}

func TestParseInTimezone(t *testing.T) {
	actual, err := ParseInTimezone("2017-07-14T02:40:00.000000+0200", Iso8601TZ, Europe_Berlin)
	assert.NoError(t, err)
	assert.Equal(
		t,
		NewDateTime(2017, 7, 14, 2, 40, 0, 0, Europe_Berlin),
		actual,
	)
}

func TestDateTime_ShiftMilliSeconds(t *testing.T) {
	dateTime := NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).ShiftMilliSeconds(10)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 1, 1, 12, 0, 0, 10000000),
		dateTime,
	)

	dateTime = NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).ShiftMilliSeconds(0)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0),
		dateTime,
	)

	dateTime = NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).ShiftMilliSeconds(-20)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 1, 1, 11, 59, 59, 980000000),
		dateTime,
	)
}

func TestDateTime_ShiftMicroSeconds(t *testing.T) {
	dateTime := NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).ShiftMicroSeconds(10)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 1, 1, 12, 0, 0, 10000),
		dateTime,
	)

	dateTime = NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).ShiftMicroSeconds(0)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0),
		dateTime,
	)

	dateTime = NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).ShiftMicroSeconds(-20)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 1, 1, 11, 59, 59, 999980000),
		dateTime,
	)
}

func TestDateTime_Shift(t *testing.T) {
	dateTime := NewUTCDateTime(2020, 1, 1, 12, 0, 1, 0).Shift(
		10,
		2,
		-3,
		-2,
		20,
		12,
		-20,
	)
	assert.Equal(
		t,
		NewUTCDateTime(2030, 2, 26, 10, 20, 12, 999999980),
		dateTime,
	)
}

func TestDateTime_Replace(t *testing.T) {
	dateTime := NewUTCDateTime(2020, 1, 1, 12, 0, 1, 0).Replace(
		2030,
		2,
		26,
		10,
		20,
		12,
		999999980,
	)
	assert.Equal(
		t,
		NewUTCDateTime(2030, 2, 26, 10, 20, 12, 999999980),
		dateTime,
	)
}

func TestDateTime_WithTimezone(t *testing.T) {

}

func TestDateTime_Timezone(t *testing.T) {
	dateTime := NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0)
	assert.Equal(t, dateTime.Timezone(), UTC)
}

func TestFromUnixTimestamp(t *testing.T) {
	timestamp := int64(1500000000)
	actual := FromUnixTimestamp(timestamp)
	assert.Equal(t, NewUTCDateTime(2017, 7, 14, 2, 40, 0, 0), actual)
}

func TestDateTime_UnixTimestamp(t *testing.T) {
	actual := NewUTCDateTime(2017, 7, 14, 2, 40, 0, 0).UnixTimestamp()
	assert.Equal(t, actual, int64(1500000000))
}

func TestDateTime_ForJson(t *testing.T) {
	actual := NewUTCDateTime(2017, 7, 14, 2, 40, 0, 0).UnixTimestamp()
	assert.Equal(t, actual, int64(1500000000))
}

func TestNewLocalDateTime(t *testing.T) {
	actual := NewLocalDateTime(2017, 7, 14, 2, 40, 0, 0)
	assert.Equal(
		t,
		actual,
		NewUTCDateTime(2017, 7, 14, 0, 40, 0, 0).InTimezone(Local()),
	)
}

func TestDateTime_IsoFormat(t *testing.T) {
	actual := NewDateTime(2017, 7, 14, 2, 40, 0, 0, Europe_Berlin).IsoFormat()
	assert.Equal(t, actual, "2017-07-14T02:40:00.000000")
}

func TestDateTime_String(t *testing.T) {
	actual := NewDateTime(2017, 7, 14, 2, 40, 0, 0, Europe_Berlin).String()
	assert.Equal(t, actual, "2017-07-14T02:40:00.000000+0200")
}

func TestDateTime_GoString(t *testing.T) {
	actual := NewDateTime(2017, 7, 14, 2, 40, 0, 0, Europe_Berlin).GoString()
	assert.Equal(t, actual, "2017-07-14T02:40:00.000000+0200")
}

func TestDateTime_IsBetween(t *testing.T) {
	actual := NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).IsBetween(
		NewUTCDateTime(2020, 1, 1, 11, 0, 0, 0),
		NewUTCDateTime(2020, 1, 1, 13, 0, 0, 0),
	)
	assert.True(t, actual)

	actual = NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).IsBetween(
		NewUTCDateTime(2020, 1, 1, 13, 0, 0, 0),
		NewUTCDateTime(2020, 1, 1, 14, 0, 0, 0),
	)
	assert.False(t, actual)
}

func TestDateTime_WeekDay(t *testing.T) {
	actual := NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).WeekDay()
	assert.Equal(t, actual, time.Wednesday)
}

func TestDateTime_IsoCalendar(t *testing.T) {
	year, month, day := NewUTCDateTime(2020, 12, 15, 12, 0, 0, 0).IsoCalendar()
	assert.Equal(t, year, 2020)
	assert.Equal(t, month, 12)
	assert.Equal(t, day, 15)
}

func TestDateTime_FloorYear(t *testing.T) {
	actual := NewUTCDateTime(2020, 12, 15, 12, 12, 49, 234).FloorYear()
	assert.Equal(t, NewUTCDateTime(2020, 1, 1, 0, 0, 0, 0), actual)
}

func TestDateTime_FloorMonth(t *testing.T) {
	actual := NewUTCDateTime(2020, 12, 15, 12, 12, 49, 234).FloorMonth()
	assert.Equal(t, NewUTCDateTime(2020, 12, 1, 0, 0, 0, 0), actual)
}

func TestDateTime_FloorDay(t *testing.T) {
	actual := NewUTCDateTime(2020, 12, 15, 12, 12, 49, 234).FloorDay()
	assert.Equal(t, NewUTCDateTime(2020, 12, 15, 0, 0, 0, 0), actual)
}

func TestDateTime_FloorHour(t *testing.T) {
	actual := NewUTCDateTime(2020, 12, 15, 12, 12, 49, 234).FloorHour()
	assert.Equal(t, NewUTCDateTime(2020, 12, 15, 12, 0, 0, 0), actual)
}

func TestDateTime_FloorMinute(t *testing.T) {
	actual := NewUTCDateTime(2020, 12, 15, 12, 12, 49, 234).FloorMinute()
	assert.Equal(t, NewUTCDateTime(2020, 12, 15, 12, 12, 0, 0), actual)
}

func TestDateTime_FloorSecond(t *testing.T) {
	actual := NewUTCDateTime(2020, 12, 15, 12, 12, 49, 234).FloorSecond()
	assert.Equal(t, NewUTCDateTime(2020, 12, 15, 12, 12, 49, 0), actual)
}

func TestDateTime_CeilYear(t *testing.T) {
	actual := NewUTCDateTime(2020, 12, 15, 12, 12, 49, 234).CeilYear()
	assert.Equal(t, NewUTCDateTime(2020, 12, 31, 23, 59, 59, 999999999), actual)
}

func TestDateTime_CeilMonth(t *testing.T) {
	actual := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).CeilMonth()
	assert.Equal(t, NewUTCDateTime(2020, 2, 29, 23, 59, 59, 999999999), actual)
}

func TestDateTime_CeilDay(t *testing.T) {
	actual := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).CeilDay()
	assert.Equal(t, NewUTCDateTime(2020, 2, 15, 23, 59, 59, 999999999), actual)
}

func TestDateTime_CeilHour(t *testing.T) {
	actual := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).CeilHour()
	assert.Equal(t, NewUTCDateTime(2020, 2, 15, 12, 59, 59, 999999999), actual)
}

func TestDateTime_CeilMinute(t *testing.T) {
	actual := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).CeilMinute()
	assert.Equal(t, NewUTCDateTime(2020, 2, 15, 12, 12, 59, 999999999), actual)
}

func TestDateTime_CeilSecond(t *testing.T) {
	actual := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).CeilSecond()
	assert.Equal(t, NewUTCDateTime(2020, 2, 15, 12, 12, 49, 999999999), actual)
}

func TestDateTime_SpanYear(t *testing.T) {
	start, end := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).SpanYear()
	assert.Equal(
		t,
		NewUTCDateTime(2020, 1, 1, 0, 0, 0, 00000000),
		start,
	)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 12, 31, 23, 59, 59, 999999999),
		end,
	)
}

func TestDateTime_SpanMonth(t *testing.T) {
	start, end := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).SpanMonth()
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 1, 0, 0, 0, 00000000),
		start,
	)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 29, 23, 59, 59, 999999999),
		end,
	)
}

func TestDateTime_SpanDay(t *testing.T) {
	start, end := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).SpanDay()
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 15, 0, 0, 0, 00000000),
		start,
	)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 15, 23, 59, 59, 999999999),
		end,
	)
}

func TestDateTime_SpanHour(t *testing.T) {
	start, end := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).SpanHour()
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 15, 12, 0, 0, 00000000),
		start,
	)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 15, 12, 59, 59, 999999999),
		end,
	)
}

func TestDateTime_SpanMinute(t *testing.T) {
	start, end := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).SpanMinute()
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 15, 12, 12, 0, 00000000),
		start,
	)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 15, 12, 12, 59, 999999999),
		end,
	)
}

func TestDateTime_SpanSecond(t *testing.T) {
	start, end := NewUTCDateTime(2020, 2, 15, 12, 12, 49, 234).SpanSecond()
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 15, 12, 12, 49, 00000000),
		start,
	)
	assert.Equal(
		t,
		NewUTCDateTime(2020, 2, 15, 12, 12, 49, 999999999),
		end,
	)
}

func TestDateTime_CTimeFormat(t *testing.T) {
	actual := NewUTCDateTime(2020, 2, 15, 12, 12, 12, 0).CTimeFormat()
	assert.Equal(
		t,
		"Sat Feb 15 12:12:12 2020",
		actual,
	)
}
