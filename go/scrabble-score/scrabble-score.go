// Package scrabble implements computing the scrabble score words.
package scrabble

import "unicode"

// Score computes scrabble scores.
func Score(word string) (scores int) {
	//word = strings.ToUpper(word)

	for _, s := range word {
		s = unicode.ToUpper(s)
		switch s {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			scores++
		case 'D', 'G':
			scores += 2
		case 'B', 'C', 'M', 'P':
			scores += 3
		case 'F', 'H', 'V', 'W', 'Y':
			scores += 4
		case 'K':
			scores += 5
		case 'J', 'X':
			scores += 8
		case 'Q', 'Z':
			scores += 10
		}
	}
	return scores
}
