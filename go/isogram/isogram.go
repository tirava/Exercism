// Package isogram implements determination if a word or phrase is an isogram.
package isogram

import "unicode"

//const (
//	space  = ' '
//	hyphen = '-'
//)

// IsIsogram returns true if string is isogram.
func IsIsogram(phrase string) bool {

	var foundRunes = map[rune]bool{}
	var sl rune

	for _, s := range phrase {
		if !unicode.IsLetter(s) {
			continue
		}
		sl = unicode.ToLower(s)
		if foundRunes[sl] {
			return false
		}
		foundRunes[sl] = true
	}

	return true
}

//func IsIsogram(input string) bool {
//	var foundRunes = map[rune]bool{}
//	input = strings.ToLower(input)
//
//	for _, val := range input {
//		if !unicode.IsLetter(val) {
//			continue
//		}
//		if foundRunes[val] {
//			return false
//		}
//		foundRunes[val] = true
//	}
//	return true
//}
