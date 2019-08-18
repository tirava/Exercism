// Package meetup implements calculating the date of meetups.
package meetup

import (
	"time"
)

// WeekSchedule type.
type WeekSchedule int

// Weekschedule constants.
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day calculates the date of the actual meetup.
func Day(ws WeekSchedule, wd time.Weekday, month time.Month, year int) int {

	t1 := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	t0 := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local)

	deltaDays := int(wd) - int(t1.Weekday())
	if deltaDays < 0 {
		deltaDays += 7
	}

	wsNum := int(ws)
	if ws == Teenth {
		wsNum = int(Second)
	}

	deltaDays += 7*wsNum + 1
	if deltaDays > t0.Day() {
		deltaDays -= 7
	}

	if ws == Teenth && deltaDays < 13 {
		deltaDays += 7
	}

	return deltaDays
}
