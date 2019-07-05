// Package pangram implements determining if a sentence is a pangram.
package pangram

import (
	"strings"
	"unicode"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func IsPangram(input string) bool {

	if len(input) < len(alphabet) {
		return false
	}

	input = strings.ToLower(input)

	for _, s := range alphabet {
		s = unicode.ToLower(s)
		if strings.Count(input, string(s)) < 1 {
			return false
		}
	}
	return true
}
