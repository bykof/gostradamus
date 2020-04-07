package gostradamus

import (
	"regexp"
	"strings"
	"time"
)

const (
	YearFull  = FormatToken("YYYY")
	YearShort = FormatToken("YY")

	MonthFull       = FormatToken("MMMM")
	MonthAbbr       = FormatToken("MMM")
	MonthZeroPadded = FormatToken("MM")
	MonthShort      = FormatToken("M")

	DayOfYearZeroPadded  = FormatToken("DDDD")
	DayOfMonthZeroPadded = FormatToken("DD")
	DayOfMonthShort      = FormatToken("D")

	DayOfWeekFullName = FormatToken("dddd")
	DayOfWeekAbbr     = FormatToken("ddd")

	TwentyFourHourZeroPadded = FormatToken("HH")
	TwelveHourZeroPadded     = FormatToken("hh")
	TwelveHour               = FormatToken("h")

	AMPMUpper = FormatToken("A")
	AMPMLower = FormatToken("a")

	MinuteZeroPadded = FormatToken("mm")
	Minute           = FormatToken("m")

	SecondZeroPadded = FormatToken("ss")
	Second           = FormatToken("s")
	MicroSecond      = FormatToken("S")

	TimezoneFullName     = FormatToken("ZZZ")
	TimezoneWithColon    = FormatToken("zz")
	TimezoneWithoutColon = FormatToken("Z")

	GoLongMonth             = goFormatToken("January")
	GoMonth                 = goFormatToken("Jan")
	GoNumMonth              = goFormatToken("1")
	GoZeroMonth             = goFormatToken("01")
	GoLongWeekDay           = goFormatToken("Monday")
	GoWeekDay               = goFormatToken("Mon")
	GoDay                   = goFormatToken("2")
	GoZeroDay               = goFormatToken("02")
	GoZeroYearDay           = goFormatToken("002")
	GoHour                  = goFormatToken("15")
	GoHour12                = goFormatToken("3")
	GoZeroHour12            = goFormatToken("03")
	GoMinute                = goFormatToken("4")
	GoZeroMinute            = goFormatToken("04")
	GoSecond                = goFormatToken("5")
	GoZeroSecond            = goFormatToken("05")
	GoLongYear              = goFormatToken("2006")
	GoYear                  = goFormatToken("06")
	GoPM                    = goFormatToken("PM")
	Gopm                    = goFormatToken("pm")
	GoTZ                    = goFormatToken("MST")
	GoMillisecond           = goFormatToken("000")
	GoMicrosecond           = goFormatToken("000000")
	GoNanoSecond            = goFormatToken("000000000")
	GoISO8601TZ             = goFormatToken("Z0700") // prints Z for UTC
	GoISO8601SecondsTZ      = goFormatToken("Z070000")
	GoISO8601ShortTZ        = goFormatToken("Z07")
	GoISO8601ColonTZ        = goFormatToken("Z07:00") // prints Z for UTC
	GoISO8601ColonSecondsTZ = goFormatToken("Z07:00:00")
	GoNumTZ                 = goFormatToken("-0700") // always numeric
	GoNumSecondsTz          = goFormatToken("-070000")
	GoNumShortTZ            = goFormatToken("-07")    // always numeric
	GoNumColonTZ            = goFormatToken("-07:00") // always numeric
	GoNumColonSecondsTZ     = goFormatToken("-07:00:00")
)

type FormatToken string
type formatTokens []FormatToken
type goFormatToken string

var (
	formatTokenMap = map[FormatToken]goFormatToken{
		YearFull:                 GoLongYear,
		YearShort:                GoYear,
		MonthFull:                GoLongMonth,
		MonthAbbr:                GoMonth,
		MonthZeroPadded:          GoZeroMonth,
		MonthShort:               GoNumMonth,
		DayOfYearZeroPadded:      GoZeroYearDay,
		DayOfMonthZeroPadded:     GoZeroDay,
		DayOfMonthShort:          GoDay,
		DayOfWeekFullName:        GoLongWeekDay,
		DayOfWeekAbbr:            GoWeekDay,
		TwentyFourHourZeroPadded: GoHour,
		TwelveHourZeroPadded:     GoZeroHour12,
		TwelveHour:               GoHour12,
		AMPMUpper:                GoPM,
		AMPMLower:                Gopm,
		MinuteZeroPadded:         GoZeroMinute,
		Minute:                   GoMinute,
		SecondZeroPadded:         GoZeroSecond,
		Second:                   GoSecond,
		MicroSecond:              GoMicrosecond,
		TimezoneFullName:         GoTZ,
		TimezoneWithColon:        GoISO8601ColonTZ,
		TimezoneWithoutColon:     GoISO8601TZ,
	}

	allFormatTokens = formatTokens{
		YearFull,
		YearShort,
		MonthFull,
		MonthAbbr,
		MonthZeroPadded,
		MonthShort,
		DayOfYearZeroPadded,
		DayOfMonthZeroPadded,
		DayOfMonthShort,
		DayOfWeekFullName,
		DayOfWeekAbbr,
		TwentyFourHourZeroPadded,
		TwelveHourZeroPadded,
		TwelveHour,
		AMPMUpper,
		AMPMLower,
		MinuteZeroPadded,
		Minute,
		SecondZeroPadded,
		Second,
		MicroSecond,
		TimezoneFullName,
		TimezoneWithColon,
		TimezoneWithoutColon,
	}
)

func (fts formatTokens) toStringSlice() []string {
	var stringSlice []string
	for _, formatToken := range fts {
		stringSlice = append(stringSlice, string(formatToken))
	}
	return stringSlice
}

func formatTokenRegex() string {
	return strings.Join(allFormatTokens.toStringSlice(), "|")
}

// translateFormat translates all mapped formats to Go specific language
func translateFormat(format string) string {
	re := regexp.MustCompile(formatTokenRegex())
	format = string(re.ReplaceAllFunc(
		[]byte(format),
		func(bytes []byte) []byte {
			if goFormatToken, ok := formatTokenMap[FormatToken(bytes)]; ok {
				return []byte(goFormatToken)
			}
			panic(FormatTokenIsNotMapped(string(bytes)))
		},
	))
	return format
}

// parseToTime parses the value with given format to a time.Time
// error if the value could not be parsed
//
// parseToTime panics if the formatToken is not mapped correctly
func parseToTime(value string, format string, timezone Timezone) (time.Time, error) {
	format = translateFormat(format)
	return time.ParseInLocation(format, value, timezone.Location())
}

// formatFromTime formats value as time.Time with given format to a string
func formatFromTime(value time.Time, format string) string {
	format = translateFormat(format)
	return value.Format(format)
}
