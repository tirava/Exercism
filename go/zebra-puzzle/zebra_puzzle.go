// Package zebra implements zebra puzzle.
package zebra

import "fmt"

// Solution base type.
type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

type house struct {
	position int
	color    string
	national string
	smoke    string
	drink    string
	animal   string
}

const num = 5

var (
	positions = [...]int{1, 2, 3, 4, 5}
	colors    = [...]string{"red", "blue", "green", "yellow", "ivory"}
	nationals = [...]string{"Englishman", "Spaniard", "Ukrainian", "Japanese", "Norwegian"}
	smokes    = [...]string{"Parliaments", "Old Gold", "Kools", "Chesterfields", "Lucky Strike"}
	drinks    = [...]string{"tea", "milk", "orange juice", "water", "coffee"}
	animals   = [...]string{"zebra", "dog", "fox", "horse", "snails"}
)

var conditions = [...]house{ // 1. There are five houses.
	{national: "Englishman", color: "red"}, // 2. The Englishman lives in the red house.
	{national: "Spaniard", animal: "dog"},  // 3. The Spaniard owns the dog.
	{color: "green", drink: "coffee"},      // 4. Coffee is drunk in the green house.
	{national: "Ukrainian", drink: "tea"},  // 5. The Ukrainian drinks tea.
	// {color: "ivory", position: x}, {color: "green", position: x+1}, // 6. The green house is immediately to the right of the ivory house.
	{smoke: "Old Gold", animal: "snails"}, // 7. The Old Gold smoker owns snails.
	{color: "yellow", smoke: "Kools"},     // 8. Kools are smoked in the yellow house.
	{position: 3, drink: "milk"},          // 9. Milk is drunk in the middle house.
	{national: "Norwegian", position: 1},  // 10. The Norwegian lives in the first house.
	// {position: y, animal: "fox"}, {position: y+1, smoke: "Chesterfields"}, // 11. The man who smokes Chesterfields lives in the house next to the man with the fox.
	// {position: z, animal: "horse"}, {position: z+1, smoke: "Kools"}, // 12. Kools are smoked in the house next to the house where the horse is kept.
	{smoke: "Lucky Strike", drink: "orange juice"}, // 13. The Lucky Strike smoker drinks orange juice.
	{national: "Japanese", smoke: "Parliaments"},   // 14. The Japanese smokes Parliaments.
	// {position: n, color: "blue"}, {national: "Norwegian", position: n + 1}, // 15. The Norwegian lives next to the blue house.
}

func generateHouses() {
	var h house

	for _, col := range colors {
		h.color = col
		for _, nat := range nationals {
			h.national = nat
			for _, smo := range smokes {
				h.smoke = smo
				for _, dri := range drinks {
					h.drink = dri
					for _, ani := range animals {
						h.animal = ani

						var hh [num]house
						for i, pos := range positions {
							h.position = pos
							hh[i] = h
							//if h.animal != "zebra" || h.drink != "water" {
							//	continue
							//}
						}

						if checkHouses(hh) {
							fmt.Println(hh)
						}
					}
				}
			}
		}
	}
}

func checkHouses(hh [num]house) bool {
	//for _, cond := range conditions {
	//	//var count, condCount int
	//	//
	//	//if cond.position != 0 {
	//	//	count++
	//	//}
	//	//if cond.color != "" {
	//	//	count++
	//	//}
	//	//if cond.national != "" {
	//	//	count++
	//	//}
	//	//if cond.smoke != "" {
	//	//	count++
	//	//}
	//	//if cond.drink != "" {
	//	//	count++
	//	//}
	//	//if cond.animal != "" {
	//	//	count++
	//	//}
	//
	//	if cond.position !=0 && cond.position  h.position {
	//		condCount++
	//	}
	//	if cond.color == h.color {
	//		condCount++
	//	}
	//	if cond.national == h.national {
	//		condCount++
	//	}
	//	if cond.smoke == h.smoke {
	//		condCount++
	//	}
	//	if cond.drink == h.drink {
	//		condCount++
	//	}
	//	if cond.animal == h.animal {
	//		condCount++
	//	}
	//
	//	fmt.Println(count, condCount, h)
	//	if count != condCount {
	//		return false
	//	}
	//}

	return true
}

// SolvePuzzle solves puzzle.
func SolvePuzzle() Solution {
	generateHouses()
	// 1. There are five houses.
	//houses := make([]house, 5)

	// 2. The Englishman lives in the red house.
	//houses[1].national = "Englishman"
	//houses[1].color = "red"

	// 3. The Spaniard owns the dog.
	//houses[3].national = "Spaniard"
	//houses[3].animal = "dog"

	// 4.

	// 5. The Ukrainian drinks tea.
	//houses[4].national = "Ukrainian"
	//houses[4].drink = "tea"

	// 9. Milk is drunk in the middle house.
	// 14. The Japanese smokes Parliaments.
	//houses[2].position = 3
	//houses[2].drink = "milk"
	//houses[2].national = "Japanese"
	//houses[2].smoke = "Parliaments"

	// 10. The Norwegian lives in the first house.
	//houses[0].position = 1
	//houses[0].national = "Norwegian"

	// 15. The Norwegian lives next to the blue house.
	//houses[34].position = 2
	//houses[34].color = "blue"

	//sort.Slice(houses, func(i, j int) bool {
	//	return houses[i].position < houses[j].position
	//})

	//for _, h := range houses {
	//	fmt.Printf("%+v\n", h)
	//}

	return Solution{}
}
