// Package wordsearch implements word search puzzles you get a square of letters
// and have to find specific words in them.
package wordsearch

import (
	"errors"
	"strings"
)

// Solve returns words coordinates.
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)

	if err := search(words, puzzle, result); err != nil {
		return nil, err
	}

	return result, nil
}

func search(words []string, puzzle []string, result map[string][2][2]int) error {
	var count int

	for i, pz := range puzzle {
		for _, w := range words {
			begin := strings.Index(pz, w)
			if begin != -1 {
				result[w] = [2][2]int{{begin, i}, {begin + len(w) - 1, i}}
				count++
			}

			begin = strings.Index(pz, reverse(w))
			if begin != -1 {
				result[w] = [2][2]int{{begin + len(w) - 1, i}, {begin, i}}
				count++
			}
		}
	}

	if count == 0 {
		return errors.New("not found")
	}

	return nil
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
