// Package yacht implements the dice game Yacht.
package yacht

var logic = map[string]func([]int, int) int{
	"yacht":  yacht,
	"ones":   anys,
	"twos":   anys,
	"threes": anys,
	"fours":  anys,
	"fives":  anys,
	"sixes":  anys,
}

var digits = map[string]int{
	"yacht":  0,
	"ones":   1,
	"twos":   2,
	"threes": 3,
	"fours":  4,
	"fives":  5,
	"sixes":  6,
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
