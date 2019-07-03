// Package isogram implements determination if a word or phrase is an isogram.
package isogram

import "strings"

const (
	space  = ' '
	hyphen = '-'
)

// IsIsogram returns true if string is isogram.
func IsIsogram(phrase string) bool {

	phrase = strings.ToLower(phrase)

	for i, s := range phrase {
		if s == space || s == hyphen {
			continue
		}

		for j := i + 1; j < len(phrase); j++ {
			//if phrase[j] == byte(s) {
			if rune(phrase[j]) == s {
				return false
			}
		}
	}

	return true
}
