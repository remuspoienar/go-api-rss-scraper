package internal

import (
	"time"
)

func ParseRFC1123ToTime(rfcTime string) time.Time {
	t, err := time.Parse(time.RFC1123, rfcTime)
	if err != nil {
		t, err = time.Parse(time.RFC3339, rfcTime)
		if err != nil {
			return time.Now().UTC()
		}
		return t.UTC()
	}
	return t.UTC()
}
