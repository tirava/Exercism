// Package meetup implements calculating the date of meetups.
package meetup

import "time"

// Weekschedule type.
type WeekSchedule int

// Weekschedule constants.
const (
	First WeekSchedule = iota + 1
	Second
	Third
	Fourth
	Last
	Teenth
) // WeekSchedule

// Day calculates the date of the actual meetup.
func Day(ws WeekSchedule, wd time.Weekday, month time.Month, year int) int {

	return 0
}
