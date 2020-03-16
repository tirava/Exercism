//// Package poker implements best hand(s) from a list of poker hands.
//package poker
//
//import (
//	"errors"
//	"strings"
//)
//
//var rank = map[string]int{"J": 100, "Q": 200, "K": 300, "A": 500}
//
//// BestHand returns best hand(s).
//func BestHand(pokers []string) ([]string, error) {
//	result := make([]string, 0)
//
//	for _, poker := range pokers {
//		cards := strings.Split(poker, " ")
//		if len(cards) != 5 {
//			return nil, errors.New("invalid cards number")
//		}
//
//		for _, card := range cards {
//			var cNum int
//			var cSuit rune
//			c := []rune(card)
//
//			if len(c) == 2 {
//				cNum, cSuit = int(c[0]-48), c[1]
//			} else if len(c) == 3 {
//				cNum = int(c[0]-48)*10 + int(c[1]-48)
//				cSuit = c[2]
//			} else {
//				return nil, errors.New("invalid card")
//			}
//
//			if cNum < 2 || (cNum > 10 && cNum != 17 && cNum != 26 && cNum != 67 && cNum != 33) {
//				return nil, errors.New("invalid card number")
//			}
//
//			if cSuit != '♡' && cSuit != '♢' && cSuit != '♤' && cSuit != '♧' {
//				return nil, errors.New("invalid card suit")
//			}
//		}
//
//		if len(pokers) == 1 {
//			return pokers, nil
//		}
//	}
//
//	return result, nil
//}

// Package poker calculate the best hand in a poker play.
package poker

import (
	"errors"
	"sort"
	"strings"
)

var (
	sign        = []rune("♡♢♧♤")
	val         = map[string]int{"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "10": 9, "J": 10, "Q": 11, "K": 12, "A": 13}
	valueString = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
)

type card struct {
	val  int
	Sign string
}

func straightColour(handsDetail []card) int64 {
	for i, handDetail := range handsDetail[1:] {
		// check same sign
		if handsDetail[0].Sign != handDetail.Sign {
			return 0
		}

		if handDetail.val != handsDetail[i].val+1 {
			return 0
		}
	}
	return int64(handsDetail[4].val) * int64(1<<(63-1))
}

func straightColourRound(handsDetail []card) (res int64) {
	res += straightColour(handsDetail)
	if handsDetail[4].val == 13 {
		res += straightColour(append([]card{card{
			val:  0,
			Sign: handsDetail[4].Sign,
		},
		}, handsDetail[:4]...))
	}
	return
}

func fourOfAKind(handsDetail []card) int64 {
	if handsDetail[1].val != handsDetail[4].val && handsDetail[0].val != handsDetail[3].val {
		return 0
	}
	return int64(handsDetail[3].val) * int64(1<<(63-4))
}

func fullHouse(handsDetail []card) int64 {
	if (handsDetail[2].val == handsDetail[4].val && handsDetail[0].val == handsDetail[1].val) ||
		(handsDetail[2].val == handsDetail[0].val && handsDetail[4].val == handsDetail[3].val) {
		return int64(handsDetail[2].val) * int64(1<<(63-8))
	}
	return 0
}

func flush(handsDetail []card) int64 {
	for _, handDetail := range handsDetail[1:] {
		// check same sign
		if handsDetail[0].Sign != handDetail.Sign {
			return 0
		}
	}
	return int64(handsDetail[4].val) * int64(1<<(63-12))
}

func straight(handsDetail []card) int64 {
	for i, handDetail := range handsDetail[1:] {
		// check same sign
		if handDetail.val != handsDetail[i].val+1 {
			return 0
		}
	}
	return int64(handsDetail[4].val) * int64(1<<(63-16))
}

func straightRound(handsDetail []card) (res int64) {
	res += straight(handsDetail)
	if handsDetail[4].val == 13 {
		res += straight(append([]card{card{
			val:  0,
			Sign: handsDetail[4].Sign,
		},
		}, handsDetail[:4]...))
	}
	return
}

func threeOfAKind(handsDetail []card) int64 {
	if handsDetail[2].val == handsDetail[4].val || handsDetail[2].val == handsDetail[0].val {
		return int64(handsDetail[2].val) * int64(1<<(63-20))
	}
	return 0
}

func pair(handsDetail []card) int64 {
	p := 0
	maxVal := 0
	sum := int64(handsDetail[0].val)
	totalSum := int64(0)
	for i, card := range handsDetail[1:] {
		if handsDetail[i].val == card.val {
			p++
			maxVal = card.val
			totalSum += int64(maxVal) * int64(1<<(63-32+p*4))
		}
		sum *= 15
		sum += int64(card.val)
	}
	return totalSum + sum
}

func calcScore(handRunes string) (score int64) {

	tmp := strings.Split(handRunes, " ")
	var e []card
	for _, t := range tmp {
		r := []rune(t)
		v, s := r[:len(r)-1], r[len(r)-1:]
		c := card{
			val:  val[string(v)],
			Sign: string(s),
		}
		e = append(e, c)
	}
	sort.Slice(e, func(i, j int) bool { return e[i].val < e[j].val })
	// start the score calculation
	// adapt looping on all the functions until one return not 0
	tail := pair(e)
	for _, f := range []func([]card) int64{straightColourRound, fourOfAKind, fullHouse, flush, straightRound, threeOfAKind, pair} {
		if s := f(e); s != 0 {
			return s + tail
		}
	}
	return
}

func containsString(array []string, key string) bool {
	for _, item := range array {
		if item == key {
			return true
		}
	}
	return false
}

func validHand(hand string) bool {
	cards := strings.Split(hand, " ")
	if len(cards) != 5 {
		return false
	}
	for _, card := range cards {
		cardRune := []rune(card)
		if len(cardRune) < 2 {
			return false
		}
		if !containsString(valueString, string(cardRune[:len(cardRune)-1])) ||
			!strings.Contains(string(sign), string(cardRune[len(cardRune)-1:])) {
			return false
		}
	}
	return true
}

// BestHand return the best hand within the set of these received as input
func BestHand(hands []string) (best []string, err error) {
	maxScore := int64(0)
	// vetting valid hands
	for _, hand := range hands {
		if !validHand(hand) {
			return nil, errors.New("invalid Hand")
		}
		score := calcScore(hand)
		switch {
		case score > maxScore:
			maxScore = score
			best = []string{hand}
		case score == maxScore:
			best = append(best, hand)
		}
	}
	return
}