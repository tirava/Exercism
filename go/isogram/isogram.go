// Package isogram implements determination if a word or phrase is an isogram.
package isogram

import (
	"strings"
)

// IsIsogram returns true if string is isogram.
func IsIsogram(phrase string) bool {

	phrase = strings.ToLower(phrase)

	for _, s := range phrase{
		s := string(s)
		if s == " " || s == "-" {
			continue
		}
		if strings.Count(phrase, s) > 1 {
			return false
		}
	}
	
	return true
}