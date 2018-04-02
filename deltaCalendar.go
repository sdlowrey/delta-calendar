/*
Package deltaCalendar calculates a date from a base date and a number of days.
 */
package deltaCalendar

import (
	"fmt"
	"time"
)

// toHours converts a number of days to a string with an "h" suffix
// The result can be used to create a time.Duration with ParseDuration.
func toHours(days int) (hours string) {
	hours = fmt.Sprintf("%dh", days * 24)
	return
}

// DeltaDate returns a time (t) that has (delta) days added to it.
// If the delta cannot be parsed as a duration, an error is returned and t
// is unchanged.
func DeltaDate(t time.Time, delta int) (time.Time, error) {
	var duration time.Duration
	var err error

	duration, err = time.ParseDuration(toHours(delta))
	if err != nil {
		return t, err
	}
	return t.Add(duration), err
}
