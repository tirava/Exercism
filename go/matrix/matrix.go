// Package matrix implements matrix of numbers, return the rows and columns of that matrix.
package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Matrix is the base type.
type Matrix [][]int

// New returns new matrix and checks them corrections.
func New(in string) (Matrix, error) {

	inR := strings.Split(in, "\n")
	m := make(Matrix, len(inR))

	i := 0
	for _, c := range inR {
		m[i] = make([]int, 0)
		inC := strings.Split(c, " ")
		for _, s := range inC {
			if s == "" {
				continue
			}
			inInt, err := strconv.Atoi(s)
			if err != nil {
				return nil, errors.New("bad matrix convert")
			}
			//fmt.Println(in, "->", inInt)
			m[i] = append(m[i], inInt)
		}
		if i > 0 && len(m[i]) != len(m[i-1]) {
			return nil, errors.New("bad matrix length")
		}
		i++
	}

	fmt.Println(m)

	return m, nil
}

// .
func (m Matrix) Set(row, col, val int) bool {

	return true
}

// Rows returns duplicate Matrix rows.
func (m Matrix) Rows() [][]int {
	duplicate := make([][]int, len(m))
	for i := range m {
		duplicate[i] = make([]int, len(m[i]))
		copy(duplicate[i], m[i])
	}
	return duplicate
}

// Cols returns duplicate Matrix columns.
func (m Matrix) Cols() [][]int {
	duplicate := make([][]int, len(m[0]))
	for j := range m[0] {
		duplicate[j] = make([]int, len(m))
		for i := range m {
			duplicate[j][i] = m[i][j]
		}
	}
	return duplicate
}
