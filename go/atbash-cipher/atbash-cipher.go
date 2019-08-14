// Package atbash implements atbash cipher, an ancient encryption system created in the Middle East.
package atbash

import (
	"strings"
	"unicode"
)

const (
	plain  = "abcdefghijklmnopqrstuvwxyz"
	cipher = "zyxwvutsrqponmlkjihgfedcba"
)

// Atbash returns ciphertext.
func Atbash(s string) (res string) {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ToLower(s)
	for i, r := range s {
		if i%5 == 0 && i != 0 {
			res += " "
		}
		if unicode.IsLetter(r) {
			indexC := strings.IndexRune(cipher, r)
			res += string(plain[indexC])
		} else {
			res += string(r)
		}
	}
	return
}
