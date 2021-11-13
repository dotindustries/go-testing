package testing

import "time"

func TimeUntilMidnight(from time.Time, timezone *time.Location) time.Duration {
	fromInTz := from.In(timezone)
	midnight := time.Date(fromInTz.Year(), fromInTz.Month(), fromInTz.Day()+1, 0, 0, 0, 0, timezone)
	return midnight.Sub(fromInTz)
}
