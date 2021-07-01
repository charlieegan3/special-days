package motheringsunday

import (
	"testing"
	"time"
)

func TestMotheringSunday(t *testing.T) {
	testCases := []struct {
		Year int
		Date time.Time
	}{
		{
			Year: 2021,
			Date: time.Date(2021, time.March, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			Year: 2022,
			Date: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
		},
		{
			Year: 2030,
			Date: time.Date(2030, time.March, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			Year: 2039,
			Date: time.Date(2039, time.March, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			Year: 1994,
			Date: time.Date(1994, time.March, 13, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range testCases {
		result, err := MotheringSunday(test.Year)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !test.Date.Equal(result) {
			t.Errorf("unexpected result: got\n%v want\n%v", result, test.Date)
		}
	}
}
