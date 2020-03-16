// Package poker implements best hand(s) from a list of poker hands.
package poker

import (
	"errors"
	"strings"
)

// BestHand returns best hand(s).
func BestHand(pokers []string) ([]string, error) {
	for _, poker := range pokers {
		cards := strings.Split(poker, " ")
		if len(cards) != 5 {
			return nil, errors.New("invalid cards number")
		}

		for _, card := range cards {
			c := []rune(card)
			if len(c) != 2 {
				return nil, errors.New("invalid card")
			}

			cNum, cSuit := c[0], c[1]
			if cNum < '2' || cNum > 10 {
				return nil, errors.New("invalid card number")
			}

			if cSuit != '♡' && cSuit != '♢' && cSuit != '♤' && cSuit != '♧' {
				return nil, errors.New("invalid card suit")
			}
		}
	}

	return nil, nil
}
