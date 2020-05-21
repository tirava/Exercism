// Package rectangles implements counting the rectangles from ASCII diagram
package rectangles

import (
	"strings"
)

func Count(input []string) int {
	var count int

	if len(input) < 2 {
		return count
	}

	//lenLine := len(input[0])
	lenInput := len(input)

	for i, line := range input {
		numPluses := strings.Count(line, "+")
		pluses := make([]int, 0, numPluses)
		for i := 0; i < numPluses; i++ {
			pp := strings.Index(line, "+")
			pluses = append(pluses, pp)
			line = strings.Replace(line, "+", "*", 1)
		}

		for p1 := 0; p1 < numPluses-1; p1++ {
			for p2 := p1 + 1; p2 < numPluses; p2++ {
				for j := i + 1; j < lenInput; j++ {
					j1, j2 := input[j][pluses[p1]], input[j][pluses[p2]]

					if j1 == '+' && j2 == '+' {
						count++
						//break
						continue
					}

					if j1 != '|' || j2 != '|' {
						break
					}
				}
			}
		}

		//for {
		//i1 := strings.Index(line, "+")
		//line = strings.Replace(line, "+", "*", 1)
		//
		//if i1 == lenLine-1 || i1 == -1 {
		//	break
		//}
		//
		//i2 := strings.Index(line, "+")
		//line = strings.Replace(line, "+", "*", 1)
		//fmt.Println(i1, i2, line)

		//for j := i + 1; j < lenInput; j++ {
		//	j1, j2 := input[j][i1], input[j][i2]
		//	//fmt.Println(string(j1), string(j2))
		//
		//	if j1 == '+' && j2 == '+' {
		//		count++
		//		break
		//	}
		//
		//	if j1 != '|' || j2 != '|' {
		//		continue
		//	}
		//
		//	//fmt.Println(string(j1), string(j2))
		//}
	}

	return count
}
