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
	return c.min2Clock((hour*60+minute)%1440 + 1440)
}

// Add minutes to Clock.
func (c Clock) Add(minutes int) Clock {
	return c.min2Clock(minutes + c.hours*60 + c.minutes)
}

// Subtract minutes from Clock.
func (c Clock) Subtract(minutes int) Clock {
	return c.min2Clock((c.hours*60+c.minutes-minutes)%1440 + 1440)
}

// String is stringer for Clock type.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

// min2Clock converts minutes to Clock type.
func (c Clock) min2Clock(minutes int) Clock {
	c.hours = minutes / 60 % 24
	c.minutes = minutes % 60
	return c
}
