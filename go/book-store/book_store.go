// Package bookstore implements trying and encourage
//more sales of different books from a popular 5 book series.
package bookstore

import "fmt"

// Cost calculates best books cost.
func Cost(books []int) int {
	var diffs [][]int

	copyBooks := make([]int, len(books))
	copy(copyBooks, books)

	for ind := 0; ; ind++ {
		diffs = append(diffs, []int{})
		for i := 1; i <= 5; i++ {
			for j, b := range copyBooks {
				if i == b {
					diffs[ind] = append(diffs[ind], b)
					copyBooks = remove(copyBooks, j)
					break
				}
			}
		}

		if len(diffs[ind]) == 0 {
			diffs = diffs[:len(diffs)-1]
			//copy(copyBooks, books)
			break
		}
	}

	for _, diff := range diffs {
		fmt.Println(diff)
	}

	fmt.Println("copy books:", copyBooks)

	fmt.Println("gen index:")
	for _, i := range genIndex() {
		fmt.Println(i)
	}

	return 0
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func genIndex() [][]int {
	var result [][]int
	d := []int{0, 5, 9, 12, 14}

	for i := 1; i <= 5; i++ {
		for k := i; k <= 5; k++ {
			result = append(result, make([]int, i))
		}
		for m := 1; m <= i; m++ {
			for j := 1; j <= 5-i+1; j++ {
				result[j+d[i-1]-1][m-1] = m
			}
		}
	}

	return result
}
