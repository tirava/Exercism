// Package transpose implements input to output text transposing.
package transpose

//
//import (
//	"strings"
//)
//
//// Transpose returns transposed text.
//func Transpose(in []string) []string {
//	max := 0
//
//	for _, s := range in {
//		if len(s) > max {
//			max = len(s)
//		}
//	}
//
//	out := make([]string, max)
//
//	for i := 0; i < max; i++ {
//		ss := strings.Builder{}
//		for _, s := range in {
//			char := "^"
//			if i < len(s) {
//				char = s[i : i+1]
//			}
//			ss.WriteString(char)
//		}
//		out[i] = ss.String()
//	}
//
//	for i := len(out) - 1; i > 0; i-- {
//		s := out[i]
//		if s[len(s)-1:] != "^" && !strings.Contains(s, "^") {
//			break
//		}
//		out[i] = strings.TrimRight(out[i], "^")
//		out[i] = strings.ReplaceAll(out[i], "^", " ")
//	}
//
//	return out
//}

// Transpose a given string array
func Transpose(m []string) []string {
	t := make([]string, maxLength(m))
	for i, row := range m {
		for j, c := range row {
			t[j] += string(c)
		}
		remMax := maxLength(m[i:])
		for j := len(row); j < remMax; j++ {
			t[j] += " "
		}
	}
	return t
}

func maxLength(m []string) (length int) {
	for _, l := range m {
		if len(l) > length {
			length = len(l)
		}
	}
	return
}
