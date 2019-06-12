// Package gigasecond implement adding 10^9 seconds
package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(1000000000 * time.Second)
	return t
}
