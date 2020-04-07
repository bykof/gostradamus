package gostradamus

import "time"

type Timezone string

func LoadLocation(timezone string) (*time.Location, error) {
	return time.LoadLocation(timezone)
}

func Local() Timezone {
	return Timezone(time.Local.String())
}

func (t Timezone) Location() *time.Location {
	location, err := LoadLocation(t.String())
	if err != nil {
		panic(err)
	}
	return location
}

func (t Timezone) String() string {
	return string(t)
}
