package motheringsunday

import (
	"fmt"
	"time"

	"github.com/vjeantet/eastertime"
)

func MotheringSunday(year int) (time.Time, error) {
	easterSunday, err := eastertime.CatholicByYear(year)
	if err != nil {
		return easterSunday, fmt.Errorf("failed to determine easter date")
	}

	motheringSunday := easterSunday.Add(-21 * 24 * time.Hour).Add(time.Hour)

	return motheringSunday.In(time.UTC), nil
}
