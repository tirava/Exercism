// Package rotationalcipher implements ROTxx, also sometimes called the Caesar cipher.
package rotationalcipher

import "unicode"

// RotationalCipher returns ROT for rot rotations.
func RotationalCipher(in string, rot int) string {

	rot %= 26
	outR := make([]rune, len(in))

	for i, s := range in {
		if !unicode.IsLetter(s) || unicode.IsSpace(s) {
			outR[i] = s
			continue
		}

		lastLetter := 'z'
		if unicode.IsUpper(s) {
			lastLetter = 'Z'
		}

		out := rune(0)
		if s+rune(rot) > lastLetter {
			out = 26
		}
		outR[i] = s + rune(rot) - out

	}

	return string(outR)
}
