// Package isogram implements determination if a word or phrase is an isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if string is isogram.
func IsIsogram(phrase string) bool {

	//phrase = strings.ToLower(phrase)
	phraseRune := []rune(phrase)
	//count := 0
	for i, r := range phraseRune {
		if !unicode.IsLower(r) {
			phraseRune[i] = unicode.ToLower(r)
		}
	}
	//unicode.strings.ToLower()
	s := ""
	for _, r := range phraseRune {
		//su := unicode.ToLower(s)
		//s := string(su)
		if r == ' ' || r == '-' { // need unicode.IsLetter
			continue
		}
		s = string(r) + s
		if strings.Count(phrase, r) > 1 {
			return false
		}
	}

	return true
}

//func countRune(rs []rune, rc rune) (count int) {
//	for _, r := range rs {
//		if r == rc {
//			count++
//			if count > 1 {
//				return
//			}
//		}
//	}
//	return
//}
