// Package wordy implements parsing and evaluating
// simple math word problems returning the answer as an integer.
package wordy

import (
	"strconv"
	"strings"
)

// Answer returns answer.
func Answer(in string) (int, bool) {

	in = strings.ReplaceAll(in, "?", "")
	in = strings.ReplaceAll(in, " by", "")
	words := strings.Split(in, " ")
	var d1, d2, result int
	var err error

	if len(words) < 3 {
		return 0, false
	}

	if len(words) == 3 {
		d1, err = strconv.Atoi(words[2])
		if err != nil {
			return 0, false
		}
		return d1, true
	}

	for i, s := range words {

		switch s {
		case "plus", "minus", "multiplied", "divided":
			if i+1 > len(words)-1 {
				return 0, false
			}
			if i+2 == len(words)-1 {
				return 0, false
			}
			d1, err = strconv.Atoi(words[i-1])
			if err != nil {
				return 0, false
			}
			d2, err = strconv.Atoi(words[i+1])
			if err != nil {
				return 0, false
			}
		}

		switch s {
		case "plus":
			if result == 0 {
				result = d1 + d2
			} else {
				result += d2
			}
		case "minus":
			if result == 0 {
				result = d1 - d2
			} else {
				result -= d2
			}
		case "multiplied":
			if result == 0 {
				result = d1 * d2
			} else {
				result *= d2
			}
		case "divided":
			if result == 0 {
				result = d1 / d2
			} else {
				result /= d2
			}
		}

	}

	if result == 0 {
		return 0, false
	}

	return result, true
}
