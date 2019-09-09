// Package wordy implements parsing and evaluating
// simple math word problems returning the answer as an integer.
package wordy

import (
	"strconv"
	"strings"
)

// Answer returns answer.
func Answer(q string) (int, bool) {

	q = strings.ReplaceAll(q, "?", "")
	q = strings.ReplaceAll(q, " by", "")

	s := strings.Split(q, " ")
	if len(s) < 3 {
		return 0, false
	}

	d1, err := strconv.Atoi(s[2]) // 1st digit
	if err != nil {
		println(d1)
		return 0, false
	}
	if len(s) < 4 {
		return d1, true
	}

	d2, err := strconv.Atoi(s[4]) // 2nd digit
	if err != nil {
		return 0, false
	}

	sign := s[3]
	switch sign {
	case "plus":
		return d1 + d2, true
	case "minus":
		return d1 - d2, true
	case "multiplied":
		return d1 * d2, true
	case "divided":
		return d1 / d2, true
	}

	return 0, false

	//return 0, false
}
