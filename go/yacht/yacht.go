// Package yacht implements the dice game Yacht.
package yacht

import "sort"

var logic = map[string]func([]int, int) int{
	"yacht":           yacht,
	"ones":            anys,
	"twos":            anys,
	"threes":          anys,
	"fours":           anys,
	"fives":           anys,
	"sixes":           anys,
	"full house":      fullHouse,
	"four of a kind":  fourOfKind,
	"little straight": straight,
	"big straight":    straight,
	"choice":          choice,
}

var digits = map[string]int{
	"ones":         1,
	"twos":         2,
	"threes":       3,
	"fours":        4,
	"fives":        5,
	"sixes":        6,
	"big straight": 1,
}

// Score returns score.
func Score(dices []int, cat string) int {
	return logic[cat](dices, digits[cat])
}

func yacht(dices []int, opts int) int {
	if dices[0] == dices[1] && dices[1] == dices[2] && dices[2] == dices[3] && dices[3] == dices[4] {
		return 50
	}

	return 0
}

func anys(dices []int, opts int) int {
	var count int

	for _, d := range dices {
		if d == opts {
			count += opts
		}
	}

	return count
}

func fullHouse(dices []int, opts int) int {
	var count int
	types := make(map[int]int)

	for _, d := range dices {
		types[d]++
		if len(types) > 2 || types[d] > 3 {
			count = 0
			break
		}

		count += d
	}

	return count
}

func fourOfKind(dices []int, opts int) int {
	var count int
	types := make(map[int]int)

	for _, d := range dices {
		types[d]++

		if types[d] == 4 {
			count = d * 4
			break
		}
	}

	return count
}

func straight(dices []int, opts int) int {
	result := 30
	sort.Ints(dices)

	if dices[0] != opts+1 {
		return 0
	}

	for i := 1; i < len(dices); i++ {
		if dices[i]+opts != dices[i-1]+1+opts {
			result = 0
			break
		}
	}

	return result
}

func choice(dices []int, opts int) int {
	var count int

	for _, d := range dices {
		count += d
	}

	return count
}
