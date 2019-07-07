// Package anagram implements finding the correct sublist.
package anagram

//import (
//	"strings"
//	"unicode"
//)
//
//// Detect returns anagram from list of candidates.
//func Detect(subject string, candidates []string) (result []string) {
//	var i int
//	var s rune
//	for _, word := range candidates {
//		wordLower := strings.ToLower(word)
//		for i, s = range subject {
//			if !strings.ContainsRune(wordLower, unicode.ToLower(s)) {
//				i = -1
//				break
//			}
//			if strings.Compare(strings.ToLower(subject), wordLower) == 0 {
//				break
//			}
//			wordLower = strings.Replace(wordLower, string(s), "_", 1)
//		}
//		if i+1 == len(word) {
//			result = append(result, word)
//		}
//	}
//	return result
//}
import (
	"sort"
	"strings"
)

func sorted(s string) string {
	runes := []rune(s)
	sort.SliceStable(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

// Detect detects anagrams of subject in a list of candidates.
func Detect(subject string, candidates []string) []string {
	anagrams := []string{}

	lowerSubject := strings.ToLower(subject)
	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		if len(candidate) != len(subject) || lowerCandidate == lowerSubject {
			continue
		}
		if sorted(lowerCandidate) == sorted(lowerSubject) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}
