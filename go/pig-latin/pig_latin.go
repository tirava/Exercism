// Package piglatin have a function Sentence to convert a sentence from english to pig latin.
package piglatin

import "strings"

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouy", r)
}

func wordConvert(s string) (res string) {
	switch {
	case (s[0] != 'y' && isVowel(rune(s[0]))) || s[:2] == "yt" || s[:2] == "xr":
		return s + "ay"
	default:
		i := 1
		for ; !isVowel(rune(s[i])); i++ {
		}
		if s[i] == 'u' && s[i-1] == 'q' {
			i++
		}
		return s[i:] + s[:i] + "ay"

	}
}

// Sentence converts a string in english into a string in pig latin.
func Sentence(s string) string {
	var tmp []string
	for _, x := range strings.Split(s, " ") {
		tmp = append(tmp, wordConvert(x))
	}
	return strings.Join(tmp, " ")
}
