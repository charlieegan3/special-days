package fathersday

import (
	"fmt"
	"time"
)

func FathersDay(country string, year int) (time.Time, error) {

	if country != "uk" {
		return time.Time{}, fmt.Errorf("country %q is not implemented", country)
	}

	sundayCount := 0
	for i := 1; i < 31; i++ {
		date := time.Date(year, time.June, i, 0, 0, 0, 0, time.UTC)
		if date.Weekday() == time.Sunday {
			sundayCount++

			if sundayCount == 3 {
				return date, nil
			}
		}
	}

	return time.Time{}, fmt.Errorf("failed to determine fathers day date")
}
