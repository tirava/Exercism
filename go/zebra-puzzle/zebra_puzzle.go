// Package zebra implements zebra puzzle.
package zebra

import (
	"fmt"
)

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
	animals   = [...]string{"snails", "fox", "dog", "horse", "zebra"}
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

func fact(n int) int {
	switch n {
	case 5:
		return 120
	case 4:
		return 24
	case 3:
		return 6
	case 2:
		return 2
	case 1, 0:
		return 1
	default:
		fmt.Println(n)
		panic("no fact!")
		//return -1
	}
}

//func permutation(index int, arr [num]int) [num]int {
//	// var n=A.length;
//	var i = index + 1
//	var res = [num]int{}
//
//	for t := 1; t <= num; t++ {
//		f := fact(num - t)
//		k := int(math.Floor((float64(i + f - 1)) / float64(f)))
//		res.push(arr.splice(k-1, 1)[0])
//		i -= (k - 1) * f
//	}
//	res.push(arr[0])
//	return res
//}

func generateHouses() {
	var hh [num]house

	var count int

	for i := range positions {
		hh[i] = house{
			position: positions[i],
			color:    colors[i],
			national: nationals[i],
			smoke:    smokes[i],
			drink:    drinks[i],
			animal:   animals[i],
		}
	}

	for i := 0; i < fact(num); i++ {
		res := make([]int, 0, num)
		source := make([]int, 0, num)
		for _, pos := range positions {
			source = append(source, pos)
		}
		for j := 0; j < num; j++ {
			p := (i / fact(num-1-j)) % len(source)
			res = append(res, source[p])
			source = append(source[:p], source[p+1:]...)
		}

		count++
		fmt.Println(count, res)
	}

	//for i := 0; i < fact(num); i++ {
	//	permutation(i, positions)
	//}

	//for _, pos := range positions {
	//	hh[i].position = pos
	//
	//	if checkHouses(hh) {
	//		count++
	//		fmt.Println(count)
	//		for _, h := range hh {
	//			fmt.Println(h)
	//		}
	//		fmt.Println("-----------------")
	//	}
	//
	//}

	fmt.Println("Houses:", count)
}

func checkHouses(hh [num]house) bool {
	var c2, c3, c4, c5, cx, cy, cz int

	for _, h := range hh {
		if h.national == "Englishman" && h.color == "red" {
			c2++
		}
		if h.national == "Spaniard" && h.animal == "dog" {
			c3++
		}
		if h.color == "green" && h.drink == "coffee" {
			c4++
		}
		if h.national == "Ukrainian" && h.drink == "tea" {
			c5++
		}

		if h.national == "Norwegian" && h.drink == "water" {
			cx++
		}
		if h.national == "Japanese" && h.animal == "zebra" { //&& h.position == 5 {
			cy++
		}

		if h.position == 1 && h.national == "Norwegian" && h.drink == "water" &&
			h.color == "yellow" && h.animal == "fox" && h.smoke == "Kools" {
			cz++
		}
	}

	//fmt.Println(c2, c3, c4, c5)
	//return c2 == 1 && c3 == 1 && c4 == 1 && c5 == 1
	//return c5 == 1 && c4 == 1 && c3 == 1
	//return cx > 0 && cy > 0 //&& c2 > 0 && c3 > 0
	return true //cz > 0 && cy > 0
}

// SolvePuzzle solves puzzle.
func SolvePuzzle() Solution {
	generateHouses()

	return Solution{}
}
