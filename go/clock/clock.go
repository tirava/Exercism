// Package clock implements a clock that handles times without dates.
package clock

import (
	"fmt"
)

// Clock type
type Clock struct {
	hours   int
	minutes int
}

// New is a "constructor".
func New(hour, minute int) (c Clock) {
	minutes := hour*60 + minute
	if minutes < 0 {
		return c.Subtract(-minutes)
	}
	return c.Add(minutes)
}

// Add minutes to Clock.
func (c Clock) Add(minutes int) Clock {
	minutes += c.hours*60 + c.minutes
	return c.normalize(minutes)
}

// Subtract minutes from Clock.
func (c Clock) Subtract(minutes int) Clock {
	minutes = (c.hours*60+c.minutes-minutes)%1440 + 1440
	return c.normalize(minutes)
}

// String is stringer for Clock type.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

// normalize does some calculations for Add & Subtract.
func (c Clock) normalize(minutes int) Clock {
	c.hours = minutes / 60
	if c.hours >= 24 {
		c.hours %= 24
	}
	c.minutes = minutes - c.hours*60
	if c.minutes >= 60 {
		c.minutes %= 60
	}
	return c
}
