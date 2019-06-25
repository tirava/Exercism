// Package isogram implements determination if a word or phrase is an isogram.
package isogram

import (
	"unicode"
)

// IsIsogram returns true if string is isogram.
func IsIsogram(phrase string) bool {

	phraseRune := []rune(phrase)

	for i, r := range phraseRune {
		phraseRune[i] = unicode.ToLower(r)
	}

	for _, s := range phraseRune {
		if !unicode.IsLetter(s) {
			continue
		}
		count := 0
		for _, c := range phraseRune {
			if s == c {
				count++
				if count > 1 {
					return false
				}
			}
		}
	}

	return true
}
