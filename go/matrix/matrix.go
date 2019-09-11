// Package matrix implements matrix of numbers, return the rows and columns of that matrix.
package matrix

// Matrix is the base type.
type Matrix [][]int

// .
func New(in string) (Matrix, error) {

	return Matrix{}, nil
}

// .
func (m Matrix) Set(row, col, val int) bool {

	return true
}

// .
func (m Matrix) Rows() [][]int {

	return nil
}

// .
func (m Matrix) Cols() [][]int {

	return nil
}
