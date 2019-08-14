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
	count := 0

	for _, r := range s {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			continue
		}
		r = unicode.ToLower(r)
		if count%5 == 0 && count != 0 {
			res += " "
		}
		count++
		if unicode.IsLetter(r) {
			indexC := strings.IndexRune(cipher, r)
			res += string(plain[indexC])
		} else {
			res += string(r)
		}
	}

	return
}
