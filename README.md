<p align="center">
    <img src="https://raw.githubusercontent.com/bykof/gostradamus/master/docs/resources/logo.png"  alt="Gostradamus logo" width="200"/>
</p>

# Gostradamus: Better DateTimes for Go

[![Gostradamus](https://img.shields.io/circleci/build/github/bykof/gostradamus)](https://app.circleci.com/pipelines/github/bykof/gostradamus)
[![Go Report Card](https://goreportcard.com/badge/github.com/bykof/gostradamus)](https://goreportcard.com/report/github.com/bykof/gostradamus)
[![codecov](https://codecov.io/gh/bykof/gostradamus/branch/master/graph/badge.svg)](https://codecov.io/gh/bykof/gostradamus)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/bykof/gostradamus)

## Introduction
Gostradamus is a Go library that offers a lightweight and human-friendly way to create, transform, format, and parse datetimes.
It uses the underlying Go `time` library and the main gostradamus' type `DateTime` can be easily converted to and from `time.Time`.

Gostradamus is named after the french pharmacist [Nostradamus](https://en.wikipedia.org/wiki/Nostradamus).
He is known for his prophecies, therefore he worked a lot with time, like Gostradamus.

## Features

✅ Easy conversion from existing `time.Time` objects to `gostradamus.DateTime` and back

✅ Timezone-aware and UTC by default

✅ Timezone conversion

✅ Generates time spans, floors, ceilings from second to year

✅ Weeks manipulation

✅ Format with common and known format tokens like `YYYY-MM-DD HH:mm:ss`


## Basic Usage

```go
package main

import "github.com/bykof/gostradamus"

func main() {
    dateTime, err := gostradamus.Parse("2017-07-14T02:40:00.000000", gostradamus.Iso8601)
    if err != nil {
        panic(err)
    }

    dateTime = dateTime.ShiftMonths(-5).ShiftDays(2)
    println(dateTime.Format("DD.MM.YYYY HH:mm:ss"))
    // 16.02.2017 02:40:00
}
```

## Table of Contents

+ [Types](#types)
+ [Conversion between time.Time and gostradamus.DateTime](#conversion-between-timetime-and-gostradamusdatetime)
+ [Creation](#creation)
+ [Timezones](#timezones)
  - [Converting](#converting)
+ [Manipulating](#manipulating)
  - [Shift](#shift)
  - [Replace](#replace)
+ [Parsing + Formatting](#parsing---formatting)
  - [Token Table](#token-table)
  - [Parsing](#parsing)
  - [Formatting](#formatting)
+ [Floor + Ceil](#floor---ceil)
  - [Floor](#floor)
  - [Ceil](#ceil)
+ [Spans](#spans)
+ [Utils](#utils)
  - [Is between](#is-between)
  - [Iso Calendar](#iso-calendar)

## Usage

This part introduces all basic features of gostradamus.
Surely there are more, just look them up in the [offical documentation](https://pkg.go.dev/github.com/bykof/gostradamus?tab=doc).

### Types

There are two types in this package, which are important to know:
```go
type DateTime time.Time
type Timezone string
```

`DateTime` contains all the creation, transforming, formatting and parsing functions.

`Timezone` is just a string type but gostradamus has all timezones defined as constants. Look [here](https://github.com/bykof/gostradamus/blob/master/timezone_constants.go).

### Conversion between time.Time and gostradamus.DateTime

You can easily convert between gostradamus.DateTime and time.Time package. 
Either with helper functions or with golang's [type conversion](https://tour.golang.org/basics/13)

```go
import "time"

// From gostradamus.DateTime to time.Time
dateTime := gostradamus.Now()
firstTime := dateTime.Time()
secondTime := time.Time(dateTime)

// From time.Time to gostradamus.DateTime
t := time.Now()
dateTime = gostradamus.DateTimeFromTime(t)
dateTime = gostradamus.DateTime(t)
```

### Creation

If you want to create a gostradamus.DateTime you have several ways:

Just create the DateTime from scratch:

```go
// Create it with a defined timezone as you know it
dateTime := gostradamus.NewDateTime(2020, 1, 1, 12, 0, 0, 0, gostradamus.EuropeBerlin)

// Create it with predefined UTC timezone
dateTime := gostradamus.NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0)

// Create it with local timzone
dateTime := gostradamus.NewLocalDateTime(2020, 1, 1, 12, 0, 0, 0)
```

Or create a DateTime from an ISO-8601 format:
```go
dateTime := gostradamus.Parse("2017-07-14T02:40:00.000000+0200", gostradamus.Iso8601)
```

Or from a custom format:
```go
dateTime := gostradamus.Parse("10.02.2010 14:59:53", "DD.MM.YYYY HH:mm:ss")
```


Or an UNIX timestamp for example:
```go 
dateTime := gostradamus.FromUnixTimestamp(1500000000)
```

Or different ways of the current datetime:
```go
// Current DateTime in local timezone
dateTime := gostradamus.Now()

// Current DateTime in UTC timezone
dateTime = gostradamus.UTCNow()

// Current DateTime in given timezone
dateTime = gostradamus.NowInTimezone(gostradamus.EuropeParis)
```

### Timezones

Feel free to use all available timezones, defined [here](https://github.com/bykof/gostradamus/blob/master/timezone_constants.go):
 
```go
gostradamus.EuropeParis // Europe/Paris
gostradamus.EuropeBerlin // Europe/Berlin
gostradamus.AmericaNewYork // America/New_York
... and many more
```

#### Converting

Convert between timezones easily:

```go
dateTime := gostradamus.NewUTC(2020, 1, 1, 12, 12, 12, 0).InTimezone(gostradamus.EuropeBerlin)
println(dateTime.String())
// 2020-02-15T13:12:12.000000+0100

dateTime = dateTime.InTimeZone(America_New_York)
println(dateTime.String())
// 2020-02-15T07:12:12.000000-0500
```


### Manipulating

#### Shift

Shifting helps you to add or subtract years, months, days, hours, minutes, seconds, milliseconds, microseconds, and nanoseconds.

To add a value use positive integer, to subtract use negative integer.

```go
dateTime := gostradamus.NewUTCDateTime(2020, 1, 1, 1, 1, 1, 1)
dateTime = dateTime.ShiftYears(10)
println(dateTime.String())
// 2030-01-01T01:01:01.000000+0000

dateTime = gostradamus.NewUTCDateTime(2020, 1, 1, 1, 1, 1, 1)
dateTime = dateTime.ShiftDays(-10)
println(dateTime.String())
// 2019-12-22T01:01:01.000000+0000

dateTime = gostradamus.NewUTCDateTime(2020, 1, 1, 1, 1, 1, 1)
dateTime = dateTime.ShiftWeeks(2)
println(dateTime.String())
// 2020-01-15T01:01:01.000000+0000

dateTime = gostradamus.NewUTCDateTime(2020, 1, 1, 1, 1, 1, 1)
dateTime = dateTime.Shift(0, 1, 10, 0, 0, 0, 0)
println(dateTime.String())
// 2020-02-11T01:01:01.000000+0000
``` 

#### Replace

Replacing values can be done easily.

```go
dateTime := gostradamus.NewUTCDateTime(2020, 1, 1, 1, 1, 1, 1)
dateTime = dateTime.ReplaceYear(2010)
println(dateTime.String())
// 2010-01-01T01:01:01.000000+0000

dateTime = gostradamus.NewUTCDateTime(2020, 1, 1, 1, 1, 1, 1)
dateTime = dateTime.ReplaceYear(2010).ReplaceMonth(2)
println(dateTime.String())
// 2010-02-01T01:01:01.000000+0000
```

### Parsing + Formatting

Parse strings easily with well-known tokens.
> Please consider that you cannot put custom tokens or custom letters into the *parsing* string 

#### Token Table

|              	| Token 	| Output                                  	|
|--------------	|-------	|-----------------------------------------	|
| Year         	| YYYY  	| 2000, 2001, 2002 … 2012, 2013           	|
|              	| YY    	| 00, 01, 02 … 12, 13                     	|
| Month        	| MMMM  	| January, February, March …              	|
|              	| MMM   	| Jan, Feb, Mar …                         	|
|              	| MM    	| 01, 02, 03 … 11, 12                     	|
|              	| M     	| 1, 2, 3 … 11, 12                        	|
| Day of Year  	| DDDD  	| 001, 002, 003 … 364, 365                	|
| Day of Month 	| DD    	| 01, 02, 03 … 30, 31                     	|
|              	| D     	| 1, 2, 3 … 30, 31                        	|
| Day of Week  	| dddd  	| Monday, Tuesday, Wednesday …            	|
|              	| ddd   	| Mon, Tue, Wed …                         	|
| Hour         	| HH    	| 00, 01, 02 … 23, 24                     	|
|              	| hh    	| 01, 02, 03 … 11, 12                     	|
|              	| h     	| 1, 2, 3 … 11, 12                        	|
| AM / PM      	| A     	| AM, PM                                  	|
|              	| a     	| am, pm                                  	|
| Minute       	| mm    	| 00, 01, 02 … 58, 59                     	|
|              	| m     	| 0, 1, 2 … 58, 59                        	|
| Second       	| ss    	| 00, 01, 02 … 58, 59                     	|
|              	| s     	| 0, 1, 2 … 58, 59                        	|
| Microsecond  	| S     	| 000000 … 999999                         	|
| Timezone     	| ZZZ   	| Asia/Baku, Europe/Warsaw, GMT           	|
|              	| zz    	| -07:00, -06:00 … +06:00, +07:00, +08, Z 	|
|              	| Z     	| -0700, -0600 … +0600, +0700, +08, Z     	|

#### Parsing

Easily parse with `Parse`:

```go
dateTime, err := gostradamus.Parse("10.02.2010 14:59:53", "DD.MM.YYYY HH:mm:ss")
println(dateTime.String())
// 2010-02-10T14:59:53.000000Z
```

You can also specify a timezone while parsing:

```go
dateTime, err := gostradamus.ParseInTimezone("10.02.2010 14:59:53", "DD.MM.YYYY HH:mm:ss", gostradamus.EuropeBerlin)
println(dateTime.String())
// 2010-02-10T14:59:53.000000+0100
```

#### Formatting

Formatting is as easy as parsing:

```go
dateTimeString := gostradamus.NewDateTime(2017, 7, 14, 2, 40, 0, 0, UTC).Format("DD.MM.YYYY Time: HH:mm:ss")
println(dateTimeString)
// 14.07.2017 Time: 02:40:00
```

### Floor + Ceil

Sometimes you need quickly the start of current day or the last day of your DateTime's month. 

That's why there is floor and ceil

#### Floor

```go
dateTimeString := gostradamus.NewDateTime(2017, 7, 14, 2, 40, 0, 0, UTC).FloorDay()
println(dateTimeString.String())
// 2017-07-14T00:00:00.000000Z

dateTimeString := gostradamus.NewDateTime(2017, 7, 14, 2, 40, 0, 0, UTC).FlorHour()
println(dateTimeString.String())
// 2017-07-14T02:00:00.000000Z
```

#### Ceil

```go
dateTimeString := gostradamus.NewDateTime(2017, 7, 14, 2, 40, 0, 0, UTC).CeilMonth()
println(dateTimeString.String())
// 2017-07-31T23:59:59.999999Z

dateTimeString := gostradamus.NewDateTime(2017, 7, 14, 2, 40, 0, 0, UTC).CeilSecond()
println(dateTimeString.String())
// 2017-07-14T02:40:00.999999Z
```

### Spans

Spans can help you to get quickly the current span of the month or the day:

```go
start, end := NewDateTime(2017, 7, 14, 2, 40, 0, 0, UTC).SpanMonth()
println(start.String())
// 2017-07-01T00:00:00.000000Z
println(end.String())
// 2017-07-31T23:59:59.999999Z

start, end = NewDateTime(2017, 7, 14, 2, 40, 0, 0, UTC).SpanDay()
println(start.String())
// 2017-07-14T00:00:00.000000Z
println(end.String())
// 2017-07-14T23:59:59.999999Z

start, end = NewDateTime(2012, 12, 12, 2, 40, 0, 0, UTC).SpanWeek()
println(start.String())
// 2012-12-10T00:00:00.000000Z
println(end.String())
// 2012-12-16T23:59:59.999999Z
```

### Utils

Here is the section for some nice helper functions that will save you some time.

#### Is between

```go
isBetween := gostradamus.NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).IsBetween(
	gostradamus.NewUTCDateTime(2020, 1, 1, 11, 0, 0, 0),
    gostradamus.NewUTCDateTime(2020, 1, 1, 13, 0, 0, 0), 
)
println(isBetween)
// true

isBetween = gostradamus.NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).IsBetween(
    gostradamus.NewUTCDateTime(2020, 1, 1, 13, 0, 0, 0),
    gostradamus.NewUTCDateTime(2020, 1, 1, 14, 0, 0, 0),
)
println(isBetween)
// false
```

#### Iso Calendar

Retrieve year, month, day directly as a 3-tuple:
 
```go
year, month, day := gostradamus.NewUTCDateTime(2020, 1, 1, 12, 0, 0, 0).IsoCalendar()
println(year, month, day)
// 2020 1 1
```







