// Package reverse implements reversing strings.
package reverse

// Reverse returns reverse string.
func Reverse(str string) string {
	strRune := []rune(str)
	lStr := len(strRune)
	for i := 0; i < lStr/2; i++ {
		strRune[i], strRune[lStr-i-1] = strRune[lStr-i-1], strRune[i]
	}
	return string(strRune)
}
