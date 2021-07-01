package fathersday

import (
	"testing"
	"time"
)

func TestFathersDay(t *testing.T) {
	testCases := []struct {
		Year int
		Date time.Time
	}{
		{
			Year: 2021,
			Date: time.Date(2021, time.June, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			Year: 2022,
			Date: time.Date(2022, time.June, 19, 0, 0, 0, 0, time.UTC),
		},
		{
			Year: 2030,
			Date: time.Date(2030, time.June, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			Year: 1994,
			Date: time.Date(1994, time.June, 19, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range testCases {
		result, err := FathersDay("uk", test.Year)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !test.Date.Equal(result) {
			t.Errorf("unexpected result: got\n%v want\n%v", result, test.Date)
		}
	}
}
