// Package railfence implements encoding and decoding for the rail fence cipher.
package railfence

import (
	"strings"
)

// Decode returns decoded string.
func Decode(in string, n int) string {
	arr := make([][]string, n)
	for i := range arr {
		arr[i] = make([]string, len(in))
	}
	indexH := 0
	indexV := 0
	indexZ := 1
	for range in {
		arr[indexV][indexH] = "?"
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

	j := 0
	inr := []rune(in)
	for _, ar := range arr {
		for i, s := range ar {
			if s == "?" {
				ar[i] = string(inr[j])
				j++
			}
		}
	}

	indexH = 0
	indexV = 0
	indexZ = 1
	out := ""
	for range in {
		out += arr[indexV][indexH]
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

	//fmt.Println(arr)

	return out
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
