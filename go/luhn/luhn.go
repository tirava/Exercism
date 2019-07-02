// Package luhn implements card validity check via Luhn formula.
package luhn

import (
	"strings"
)

// Valid returns true if card number is valid.
func Valid(number string) bool {

	number = strings.ReplaceAll(number, " ", "")
	if len(number) < 2 {
		return false
	}

	lenNumber := len(number)
	numbers := make([]int, lenNumber)
	var numInt, i, d int

	for i = 0; i < lenNumber; i++ {
		numInt = int(number[lenNumber-1-i]) - 48
		if numInt > 9 {
			return false
		}
		if i%2 == 0 {
			numbers[i] = numInt
		} else {
			d = numInt * 2
			if d > 9 {
				d -= 9
			}
			numbers[i] = d
		}
	}

	var sum int
	for _, n := range numbers {
		sum += n
	}
	if sum%10 == 0 {
		return true
	}

	return false
}
