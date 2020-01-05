// Package pascal implements computing Pascal's triangle up to a given number of rows.
package pascal

// Triangle returns Pascal's triangle given size.
func Triangle(n int) [][]int {
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i][j] = 1
				continue
			}
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}
