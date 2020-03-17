// Package spiralmatrix implements a square matrix of numbers in spiral order.
package spiralmatrix

// SpiralMatrix returns matrix.
func SpiralMatrix(in int) [][]int {
	if in <= 0 {
		return [][]int{}
	}

	if in == 1 {
		return [][]int{{1}}
	}

	result := make([][]int, in)
	for i := 0; i < in; i++ {
		result[i] = make([]int, in)
	}

	count := 1
	for index := 0; index < in; index++ {
		for i := index; i < len(result[0])-index*2; i++ {
			result[index][i] = count
			count++
		}

		for j := index + 1; j < len(result)-index*2; j++ {
			result[j][in-1-index] = count
			count++
		}

		for i := len(result[0]) - 2 - index; i >= index; i-- {
			result[in-index-1][i] = count
			count++
		}

		for j := in - 2 - index; j > index; j-- {
			result[j][index] = count
			count++
		}
	}

	return result
}
