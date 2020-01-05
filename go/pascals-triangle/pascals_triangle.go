// Package pascal implements computing Pascal's triangle up to a given number of rows.
package pascal

// Triangle returns Pascal's triangle given size.
func Triangle(n int) [][]int {
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, i)
	}

	return result
}
