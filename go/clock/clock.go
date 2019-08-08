// Package clock implements a clock that handles times without dates.
package clock

import (
	"fmt"
)

type Clock int

const (
	minsInHour = 60
	hoursInDay = 24
	minsInDay  = minsInHour * hoursInDay
)

// New is a "constructor".
func New(hour, minute int) Clock {
	return Clock((hour*minsInHour+minute)%minsInDay+minsInDay) % minsInDay
}

// Add minutes to Clock.
func (c Clock) Add(minutes int) Clock {
	return c + Clock(minutes)
}

// Subtract minutes from Clock.
func (c Clock) Subtract(minutes int) Clock {
	return (c - Clock(minutes)%minsInDay + minsInDay) % minsInDay
}

// String is stringer for Clock type.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/minsInHour%hoursInDay, c%minsInHour)
}
