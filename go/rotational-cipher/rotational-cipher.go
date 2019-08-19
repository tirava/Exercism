// Package rotationalcipher implements ROTxx, also sometimes called the Caesar cipher.
package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher returns ROT for rot rotations.
func RotationalCipher(in string, rot int) string {

	var (
		outR strings.Builder
		resR rune
		rotr = rune(rot % 26)
	)

	outR.Grow(len(in))

	for _, s := range in {
		if unicode.IsLetter(s) {
			lastLetter := 'z'
			if unicode.IsUpper(s) {
				lastLetter = 'Z'
			}

			var out rune = 0
			if s+rotr > lastLetter {
				out = 26
			}
			resR = s + rotr - out
		} else {
			resR = s
		}

		outR.WriteRune(resR)
	}

	return outR.String()
}
