// Package zebra implements zebra puzzle.
package zebra

import (
	"fmt"
	"sort"
)

// Solution base type.
type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

type house struct {
	color    string
	position int
	national string
	smoke    string
	drink    string
	animal   string
}

// SolvePuzzle solves puzzle.
func SolvePuzzle() Solution {
	// 1. There are five houses.
	houses := make([]house, 5)

	// 2. The Englishman lives in the red house.
	houses[1].national = "Englishman"
	houses[1].color = "red"

	// 3. The Spaniard owns the dog.
	houses[3].national = "Spaniard"
	houses[3].animal = "dog"

	// 4.

	// 5. The Ukrainian drinks tea.
	houses[4].national = "Ukrainian"
	houses[4].drink = "tea"

	// 9. Milk is drunk in the middle house.
	// 14. The Japanese smokes Parliaments.
	houses[2].position = 3
	houses[2].drink = "milk"
	houses[2].national = "Japanese"
	houses[2].smoke = "Parliaments"

	// 10. The Norwegian lives in the first house.
	houses[0].position = 1
	houses[0].national = "Norwegian"

	// 15. The Norwegian lives next to the blue house.
	//houses[34].position = 2
	//houses[34].color = "blue"

	sort.Slice(houses, func(i, j int) bool {
		return houses[i].position < houses[j].position
	})

	for _, h := range houses {
		fmt.Printf("%+v\n", h)
	}

	return Solution{}
}
