// Package poker implements best hand(s) from a list of poker hands.
package poker

import (
	"errors"
	"strings"
)

var rank = map[string]int{"J": 100, "Q": 200, "K": 300, "A": 500}

// BestHand returns best hand(s).
func BestHand(pokers []string) ([]string, error) {
	result := make([]string, 0)

	for _, poker := range pokers {
		cards := strings.Split(poker, " ")
		if len(cards) != 5 {
			return nil, errors.New("invalid cards number")
		}

		for _, card := range cards {
			var cNum int
			var cSuit rune
			c := []rune(card)

			if len(c) == 2 {
				cNum, cSuit = int(c[0]-48), c[1]
			} else if len(c) == 3 {
				cNum = int(c[0]-48)*10 + int(c[1]-48)
				cSuit = c[2]
			} else {
				return nil, errors.New("invalid card")
			}

			if cNum < 2 || (cNum > 10 && cNum != 17 && cNum != 26 && cNum != 67 && cNum != 33) {
				return nil, errors.New("invalid card number")
			}

			if cSuit != '♡' && cSuit != '♢' && cSuit != '♤' && cSuit != '♧' {
				return nil, errors.New("invalid card suit")
			}
		}

		if len(pokers) == 1 {
			return pokers, nil
		}
	}

	return result, nil
}
