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

	search1 := search(words, puzzle, result, false)
	var search2 bool
	//fmt.Println(puzzle)
	if len(words) > 1 {
		puzzleReverse := reversePuzzle(puzzle)
		//fmt.Println(puzzleReverse)
		search2 = search(words, puzzleReverse, result, true)
	}

	if !search1 && !search2 {
		return nil, errors.New("not found")
	}

	return result, nil
}

func reversePuzzle(words []string) []string {
	newWords := make([]string, len(words[0]))

	newWord := strings.Builder{}
	for i := 0; i < len(words[0]); i++ {
		for j := 0; j < len(words); j++ {
			newWord.WriteByte(words[j][i])
		}

		newWords[i] = newWord.String()
		newWord.Reset()
	}

	return newWords
}

func search(words []string, puzzle []string, result map[string][2][2]int, reversed bool) bool {
	var count int

	for i, pz := range puzzle {
		for _, w := range words {
			begin := strings.Index(pz, w)
			if begin != -1 {
				if !reversed {
					result[w] = [2][2]int{{begin, i}, {begin + len(w) - 1, i}}
				} else {
					result[w] = [2][2]int{{i, begin}, {i, begin + len(w) - 1}}
				}
				count++
			}

			begin = strings.Index(pz, reverse(w))
			if begin != -1 {
				if !reversed {
					result[w] = [2][2]int{{begin + len(w) - 1, i}, {begin, i}}
				} else {
					result[w] = [2][2]int{{i, begin + len(w) - 1}, {i, begin}}
				}
				count++
			}
		}
	}

	return count != 0
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
