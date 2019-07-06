// Package wordcount implements words counting.
package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

// WordCount counts every word in phrase.
//func WordCount(input string) (counts Frequency) {
//
//	counts = make(Frequency)
//
//	input = strings.ToLower(input)
//	reg, _ := regexp.Compile("[^a-z0-9']+")
//	input = reg.ReplaceAllString(input, "_")
//
//	words := strings.Split(input, "_")
//
//	for _, word := range words {
//		if word == "" {
//			continue
//		}
//		word = strings.Trim(word, "'")
//		counts[word]++
//	}
//	return counts
//}

//func WordCount(phrase string) Frequency {
//
//	phrase = strings.ToLower(phrase)
//
//	var result = make(Frequency)
//
//	f := func(c rune) bool {
//		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
//	}
//
//	words := strings.FieldsFunc(phrase, f)
//
//	for _, word := range words {
//		word = strings.Trim(word, "'")
//		result[word]++
//	}
//
//	return result
//}

//type Frequency map[string]int

var reg = regexp.MustCompile("[a-z0-9]+('[a-z0-9])*")

// WordCount return the frequencies of a word in a string
func WordCount(s string) Frequency {
	s = strings.ToLower(s)
	result := Frequency{}
	words := reg.FindAllString(s, -1)
	for _, c := range words {
		result[c]++
	}
	return result
}
