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

		for i := index; i < in-index; i++ {
			result[index][i] = count
			count++
		}

		for j := index + 1; j < in-index; j++ {
			result[j][in-1-index] = count
			count++
		}

		for i := in - 2 - index; i >= index; i-- {
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

//package spiralmatrix
//
//// SpiralMatrix returns a spiral matrix of size n
//func SpiralMatrix(n int) [][]int {
//	m := make([][]int, n)
//	for y := 0; y < n; y++ {
//		m[y] = make([]int, n)
//	}
//
//	x, y := 0, 0
//	dx, dy := 1, 0
//	for count := 1; count <= n*n; count++ {
//		m[y][x] = count
//		xx, yy := x+dx, y+dy
//		if xx < 0 || xx >= n || yy < 0 || yy >= n || m[yy][xx] != 0 {
//			dx, dy = -dy, dx
//		}
//		x, y = x+dx, y+dy
//	}
//
//	return m
//}
