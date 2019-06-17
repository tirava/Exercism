// Package raindrops implements converting a number to a string, in depend on the number's factors.
package raindrops

import "strconv"

// Convert need for converting a number to a string
// the contents of which depend on the number's factors.
func Convert(input int) string {
	var pling, plang, plong string

	if input%3 == 0 {
		pling = "Pling"
	}
	if input%5 == 0 {
		plang = "Plang"
	}
	if input%7 == 0 {
		plong = "Plong"
	}

	s := pling + plang + plong
	if len(s) == 0 {
		s = strconv.Itoa(input)
	}

	return s
}
