// Package rectangles implements counting the rectangles from ASCII diagram
package rectangles

import (
	"strings"
)

// Count counts rectangles.
func Count(input []string) int {
	var count int

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
						continue
					}

					if (j1 == '+' && j2 == '|') || (j1 == '|' && j2 == '+') {
						continue
					}

					if j1 != '|' || j2 != '|' {
						break
					}
				}
			}
		}
	}

	return count
}
