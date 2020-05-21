// Package rectangles implements counting the rectangles from ASCII diagram
package rectangles

import (
	"fmt"
	"strings"
)

func Count(input []string) int {
	var count int

	if len(input) < 2 {
		return count
	}

	lenLine := len(input[0])
	lenInput := len(input)

	for i, line := range input {
		for {
			i1 := strings.Index(line, "+")
			line = strings.Replace(line, "+", "*", 1)

			if i1 == lenLine-1 || i1 == -1 {
				break
			}

			i2 := strings.Index(line, "+")
			line = strings.Replace(line, "+", "*", 1)
			//fmt.Println(i1, i2, line)

			for j := i + 1; j < lenInput; j++ {
				j1, j2 := input[j][i1], input[j][i2]
				fmt.Println(string(j1), string(j2))

				if j1 == '+' && j2 == '+' {
					count++
					break
				}

				if j1 != '|' || j2 != '|' {
					continue
				}

				//fmt.Println(string(j1), string(j2))
			}
		}

	}

	return count
}
