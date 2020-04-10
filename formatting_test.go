package gostradamus

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseToTime(t *testing.T) {
	actualResult, err := parseToTime("2019-01-01 12:12:12", "YYYY-MM-DD HH:mm:ss", UTC)
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2019, 1, 1, 12, 12, 12, 0, time.UTC), actualResult)
}

func TestParseToTime_Error(t *testing.T) {
	actualResult, err := parseToTime("2019-01-01 12:12:12", "YYYY-MM-DD testHH:mm:ss", UTC)
	assert.Error(t, err)
	assert.Equal(t, time.Time{}, actualResult)
}

func TestFormatFromTime(t *testing.T) {
	actualResult := formatFromTime(
		time.Date(2011, 4, 5, 15, 7, 8, 9, time.UTC),
		"YYYY YY MMMM MMM MM M DDDD DD D dddd ddd HH hh h A a mm m ss s S ZZZ zz Z",
	)
	assert.Equal(
		t,
		"2011 11 April Apr 04 4 095 05 5 Tuesday Tue 15 03 3 PM pm 07 7 08 8 000000 UTC Z Z",
		actualResult,
	)

	location, err := time.LoadLocation("Japan")
	if err != nil {
		panic(err)
	}

	date := time.Date(2011, 4, 5, 15, 7, 8, 9, location)
	actualResult = formatFromTime(
		date,
		"ZZZ zz Z",
	)

	assert.Equal(
		t,
		"JST +09:00 +0900",
		actualResult,
	)
}
