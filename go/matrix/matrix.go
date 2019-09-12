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

// .
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

// .
func (m Matrix) Rows() [][]int {

	return nil
}

// .
func (m Matrix) Cols() [][]int {

	return nil
}
