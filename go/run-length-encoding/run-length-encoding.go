// Package encode implements RLE encode/decode.
package encode

import (
	"strconv"
	"unicode"
)

type runCounts map[rune]int

// RunLengthEncode returns RLE encoding.
func RunLengthEncode(in string) (out string) {
	if len(in) < 2 {
		return ""
	}

	inRune := []rune(in)
	counts := make(runCounts)

	counts[inRune[0]]++

	for i := 1; i < len(inRune); i++ {
		if inRune[i] == inRune[i-1] {
			counts[inRune[i]]++
			if i == len(inRune)-1 {
				out += strconv.Itoa(counts[inRune[i-1]]) + string(inRune[i])
			}
		} else {
			if counts[inRune[i-1]] != 1 {
				out += strconv.Itoa(counts[inRune[i-1]])
			}
			out += string(inRune[i-1])
			counts = make(runCounts)
			counts[inRune[i]]++
			if i == len(inRune)-1 {
				out += string(inRune[i])
			}
		}
	}

	return out
}

// RunLengthDecode returns RLE decoding.
func RunLengthDecode(in string) (out string) {
	if len(in) < 2 {
		return ""
	}

	inRune := []rune(in)
	ds, dd, j := "", 0, 0

	//LOOP:
	for i := 0; i < len(inRune); i++ {
		if unicode.IsLetter(inRune[i]) || unicode.IsSpace(inRune[i]) {
			out += string(inRune[i])
		} else if unicode.IsDigit(inRune[i]) {
			ds = ""
			for j = i; j < len(inRune); j++ { // get all digits
				//i++
				dd = int(inRune[j] - 48)
				if dd >= 0 && dd <= 9 {
					ds += string(inRune[j])
					i++
				} else {
					dd, _ = strconv.Atoi(ds)
					for k := 0; k < dd; k++ {
						out += string(inRune[i])
					}
					break
				}
			}

		} else {
			return ""
		}
	}

	return out
}
