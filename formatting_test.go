package gostradamus

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseToTime(t *testing.T) {
	actualResult, err := ParseToTime("2019-01-01 12:12:12", "YYYY-MM-DD HH:mm:ss", UTC)
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2019, 1, 1, 12, 12, 12, 0, time.UTC), actualResult)
}

func TestParseToTime_Error(t *testing.T) {
	actualResult, err := ParseToTime("2019-01-01 12:12:12", "YYYY-MM-DD testHH:mm:ss", UTC)
	assert.Error(t, err)
	assert.Equal(t, time.Time{}, actualResult)
}

func TestFormatFromTime(t *testing.T) {
	actualResult := FormatFromTime(
		time.Date(2019, 1, 1, 12, 12, 12, 0, time.UTC),
		"YYYY-MM-DD HH:mm:ss",
	)
	assert.Equal(t, actualResult, "2019-01-01 12:12:12")
}
