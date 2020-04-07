package gostradamus

const (
	// Iso8601 format
	//
	//     Example: 2012-12-12T12:12:12.000000
	//
	Iso8601 = "YYYY-MM-DDTHH:mm:ss.S"

	// Iso8601TZ format with timezone
	//
	//     Example with UTC: 2012-12-12T12:12:12.000000Z
	//     Example with timezone: 2012-12-12T12:12:12.000000+0200
	//
	Iso8601TZ = "YYYY-MM-DDTHH:mm:ss.SZ"

	// CTime format with ctime standard
	//
	//     Example: Sat Jan 19 18:26:50 2019
	//
	CTime = "ddd MMM DD HH:mm:ss YYYY"
)
