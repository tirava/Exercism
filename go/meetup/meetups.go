// Package meetup implements calculating the date of meetups.
package meetup

import (
	"fmt"
	"time"
)

// Weekschedule type.
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

	p := fmt.Println

	t1 := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	deltaDays := int(wd) - int(t1.Weekday())
	if deltaDays < 0 {
		deltaDays += 7
	}

	deltaDays += 7*int(ws) + 1

	//if t1.Weekday() != 0 {
	//	deltaDays += 7-int(t1.Weekday())
	//}

	//p(int(ws), int(wd), int(t1.Weekday()))
	//if t1.Weekday() == time.Monday {
	//	deltaDays -= 7
	//}

	//t1.Add(time.Hour * 24 * time.Duration(deltaDays))

	p(t1.Weekday(), wd, deltaDays)

	return deltaDays
}
