// Package clock implements a clock that handles times without dates.
package clock

import "strconv"

// Clock type
type Clock string

// New is a "constructor".
func New(hour, minute int) (c Clock) {
	minutes := hour*60 + minute
	return c.Add(minutes)
}

// Add minutes to Clock.
func (c Clock) Add(minutes int) Clock {
	hours := minutes / 60
	if hours >= 24 {
		hours %= 24
	}
	minutes = minutes - hours*60
	if minutes >= 60 {
		minutes %= 60
	}
	h := strconv.Itoa(hours)
	m := strconv.Itoa(minutes)
	if hours < 10 {
		h = "0" + h
	}
	if minutes < 10 {
		m = "0" + m
	}
	return Clock(h + ":" + m)
}

// Subtract minutes to Clock.
func (c Clock) Subtract(minutes int) Clock {

	return ""
}

func (c Clock) String() string {
	return string(c)
}
