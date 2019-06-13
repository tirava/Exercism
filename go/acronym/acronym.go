// Package acronym implements converting a phrase to its acronym.
package acronym

import "strings"

var replacer = strings.NewReplacer("-", " ", "_", " ")

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	var acr string
	s1 := replacer.Replace(s)
	words := strings.Split(s1 , " ")
	for _, word := range words {
		if word == "" {
			continue
		}
		letter := strings.ToUpper(word[0:1])
		acr += letter
	}
	return acr
}
