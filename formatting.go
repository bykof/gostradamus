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

	GoLongMonth             = GoFormatToken("January")
	GoMonth                 = GoFormatToken("Jan")
	GoNumMonth              = GoFormatToken("1")
	GoZeroMonth             = GoFormatToken("01")
	GoLongWeekDay           = GoFormatToken("Monday")
	GoWeekDay               = GoFormatToken("Mon")
	GoDay                   = GoFormatToken("2")
	GoZeroDay               = GoFormatToken("02")
	GoZeroYearDay           = GoFormatToken("002")
	GoHour                  = GoFormatToken("15")
	GoHour12                = GoFormatToken("3")
	GoZeroHour12            = GoFormatToken("03")
	GoMinute                = GoFormatToken("4")
	GoZeroMinute            = GoFormatToken("04")
	GoSecond                = GoFormatToken("5")
	GoZeroSecond            = GoFormatToken("05")
	GoLongYear              = GoFormatToken("2006")
	GoYear                  = GoFormatToken("06")
	GoPM                    = GoFormatToken("PM")
	Gopm                    = GoFormatToken("pm")
	GoTZ                    = GoFormatToken("MST")
	GoMillisecond           = GoFormatToken("000")
	GoMicrosecond           = GoFormatToken("000000")
	GoNanoSecond            = GoFormatToken("000000000")
	GoISO8601TZ             = GoFormatToken("Z0700") // prints Z for UTC
	GoISO8601SecondsTZ      = GoFormatToken("Z070000")
	GoISO8601ShortTZ        = GoFormatToken("Z07")
	GoISO8601ColonTZ        = GoFormatToken("Z07:00") // prints Z for UTC
	GoISO8601ColonSecondsTZ = GoFormatToken("Z07:00:00")
	GoNumTZ                 = GoFormatToken("-0700") // always numeric
	GoNumSecondsTz          = GoFormatToken("-070000")
	GoNumShortTZ            = GoFormatToken("-07")    // always numeric
	GoNumColonTZ            = GoFormatToken("-07:00") // always numeric
	GoNumColonSecondsTZ     = GoFormatToken("-07:00:00")
)

type FormatToken string
type FormatTokens []FormatToken
type GoFormatToken string
type GoFormatTokens []GoFormatToken

var (
	FormatTokenMap = map[FormatToken]GoFormatToken{
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

	AllFormatTokens = FormatTokens{
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

func (fts FormatTokens) ToStringSlice() []string {
	var stringSlice []string
	for _, formatToken := range fts {
		stringSlice = append(stringSlice, string(formatToken))
	}
	return stringSlice
}

func FormatTokenRegex() string {
	return strings.Join(AllFormatTokens.ToStringSlice(), "|")
}

func translateFormat(format string) string {
	re := regexp.MustCompile(FormatTokenRegex())
	format = string(re.ReplaceAllFunc(
		[]byte(format),
		func(bytes []byte) []byte {
			if goFormatToken, ok := FormatTokenMap[FormatToken(bytes)]; ok {
				return []byte(goFormatToken)
			}
			panic(FormatTokenIsNotMapped(string(bytes)))
		},
	))
	return format
}

// ParseToTime parses the value with given format to a time.Time
// error if the value could not be parsed
//
// ParseToTime panics if the formatToken is not mapped correctly
func ParseToTime(value string, format string, timezone Timezone) (time.Time, error) {
	format = translateFormat(format)
	return time.ParseInLocation(format, value, timezone.Location())
}

//
func FormatFromTime(value time.Time, format string) string {
	format = translateFormat(format)
	return value.Format(format)
}
