// Package say implements the int to string conversion.
package say

import (
	"strings"
)

var numToString = map[int64]string{
	0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
	11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty",
	30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety", 100: "hundred",
}

var thousands = []string{"", "thousand", "million", "billion"}

// Say converts a number into a string.
func Say(x int64) (s string, ok bool) {

	ok = true
	if x < 0 || x > 999999999999 {
		ok = false
		return
	}
	if x == 0 {
		return numToString[0], true
	}

	for _, item := range thousands {
		var tmp []string
		local := x % 1000

		if local >= 100 {
			tmp = append(tmp, numToString[local/100]+" "+numToString[100])
			local %= 100
		}
		if local > 20 {
			tmp = append(tmp, numToString[(local/10)*10]+"-"+numToString[local%10])
		}
		if local > 0 && local < 21 {
			tmp = append([]string{numToString[local]}, tmp...)
		}

		if len(item) > 0 && len(tmp) > 0 {
			tmp = append(tmp, item)
		}
		if len(s) != 0 {
			s = " " + s
		}
		s = strings.Join(tmp, " ") + s
		x /= 1000
		if x == 0 {
			break
		}
	}
	return
}
