// Package encode implements RLE encode/decode.
package encode

import (
	"fmt"
	"strconv"
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
	//fmt.Println("counts[0]:", counts)

	for i := 1; i < len(inRune); i++ {
		//fmt.Println(string(inRune[i]))
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
			//fmt.Println("out:", out)
		}
		//println(i, string(inRune[i]))
	}

	fmt.Println(string(inRune), out)
	return out
}

// RunLengthDecode returns RLE decoding.
func RunLengthDecode(input string) string {
	if len(input) < 2 {
		return ""
	}

	return ""
}
