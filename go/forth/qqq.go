// Package forth implements ... .
package forth

func isMagic(square [][]int) bool {
	var magic int

	for _, v := range square[0] {
		magic += v
	}

	lenSquare := len(square[0])

	// horizontal
	for i := 0; i < lenSquare; i++ {
		var sum int

		for _, v := range square[i] {
			sum += v
		}

		if sum != magic {
			return false
		}
	}

	// vertical
	for j := 0; j < lenSquare; j++ {
		var sum int

		for i := 0; i < lenSquare; i++ {
			sum += square[i][j]
		}

		if sum != magic {
			return false
		}
	}

	var sum int

	// diagonal 1
	for i, j := 0, 0; i < lenSquare; i, j = i+1, j+1 {
		sum += square[i][j]
	}

	if sum != magic {
		return false
	}

	sum = 0

	// diagonal 2
	for i, j := lenSquare-1, lenSquare-1; i >= 0; i, j = i-1, j-1 {
		sum += square[i][j]
	}

	return sum == magic
}
