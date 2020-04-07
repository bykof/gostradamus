package gostradamus

import "time"

// Timezone is a string type which can translate from and to time.Location
type Timezone string

// LoadLocation returns the time.Locatior or an error of a timezone
func LoadLocation(timezone string) (*time.Location, error) {
	return time.LoadLocation(timezone)
}

// Local is a wrapper for "Local" as a Timezone
func Local() Timezone {
	return Timezone(time.Local.String())
}

// Location returns the Location of current Timezone
//
// Location panics if current Timzeone does not exist
func (t Timezone) Location() *time.Location {
	location, err := LoadLocation(t.String())
	if err != nil {
		panic(err)
	}
	return location
}

// String returns Timezone as string
// Example: "Europe/Berlin"
func (t Timezone) String() string {
	return string(t)
}
