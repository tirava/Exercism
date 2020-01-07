// Package railfence implements encoding and decoding for the rail fence cipher.
package railfence

import (
	"strings"
)

// Decode returns decoded string.
func Decode(in string, n int) string {

	return ""
}

// Encode returns encoded string.
func Encode(in string, n int) string {
	arr := make([][]string, n)
	for i := range arr {
		arr[i] = make([]string, len(in))
	}
	indexH := 0
	indexV := 0
	indexZ := 1
	for _, s := range in {
		arr[indexV][indexH] = string(s)
		indexH++
		if indexH >= len(in) {
			indexH = 0
		}
		indexV += indexZ
		if indexV >= n || indexV < 0 {
			indexV -= indexZ
			indexZ = -indexZ
			indexV += indexZ
		}
	}

	out := ""
	for _, ar := range arr {
		s := strings.Join(ar, "")
		out += s
	}

	return out
}
