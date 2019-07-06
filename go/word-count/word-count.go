// Package wordcount implements words counting.
package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

// WordCount counts every word in phrase.
func WordCount(input string) (counts Frequency) {

	counts = make(Frequency)

	input = strings.ToLower(input)
	reg, _ := regexp.Compile("[^a-z0-9']+")
	input = reg.ReplaceAllString(input, "_")

	words := strings.Split(input, "_")

	for _, word := range words {
		if word == "" {
			continue
		}
		word = strings.Trim(word, "'")
		counts[word]++
	}
	return counts
}
