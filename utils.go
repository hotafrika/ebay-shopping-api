package shopping

import (
	"time"
	"unicode"
)

var UTC, _ = time.LoadLocation("UTC")

// FromEbayDateTime converts eBay datetime format to time.Time
// By default eBay date-time values are recorded in Universal Coordinated Time (UTC)
func FromEbayDateTime(ebayDT string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02T15:04:05.000Z", ebayDT, UTC)
}

// ToEbayDateTime converts given time to eBay format
// Given datetime has to be casted to UTC location.
func ToEbayDateTime(datetime time.Time) string {
	return datetime.Format("2006-01-02T15:04:05.000Z")
}

// FromEbayDuration converts eBay duration to Golang duration
// eBay format is PnYnMnDTnHnMnS (e.g., P2DT23H32M51S)
func FromEbayDuration(ebayDuration string) time.Duration {
	d := time.Second * 0
	var cd int
	for _, b := range ebayDuration {
		if unicode.IsNumber(b) {
			cd = cd*10 + int(b) - '0' // convert rune to int
			continue
		}
		switch b {
		case 'D':
			d = d + time.Hour*24*time.Duration(cd)
		case 'H':
			d = d + time.Hour*time.Duration(cd)
		case 'M':
			d = d + time.Minute*time.Duration(cd)
		case 'S':
			d = d + time.Second*time.Duration(cd)
		case 'P', 'T':
			// empty
		default:
			continue
		}
		cd = 0
	}

	return d
}
