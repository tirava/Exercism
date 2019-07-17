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
	var sum int

	for i := 0; i < lenNumber; i++ {
		numInt := int(number[lenNumber-1-i]) - '0'
		if numInt > 9 {
			return false
		}
		if i%2 == 0 {
			sum += numInt
		} else {
			d := numInt * 2
			if d > 9 {
				d -= 9
			}
			sum += d
		}
	}

	return sum%10 == 0
}
