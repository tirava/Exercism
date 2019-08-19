// Package rotationalcipher implements ROTxx, also sometimes called the Caesar cipher.
package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher returns ROT for rot rotations.
func RotationalCipher(in string, rot int) string {

	rotr := rune(rot % 26)
	var outR strings.Builder
	outR.Grow(len(in))

	for _, s := range in {
		if !unicode.IsLetter(s) || unicode.IsSpace(s) {
			outR.WriteRune(s)
			continue
		}

		lastLetter := 'z'
		if unicode.IsUpper(s) {
			lastLetter = 'Z'
		}

		out := rune(0)
		if s+rotr > lastLetter {
			out = 26
		}
		outR.WriteRune(s + rotr - out)

	}
	return outR.String()
}
