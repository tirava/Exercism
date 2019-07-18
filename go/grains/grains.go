// Package grains implements calculating the number of grains of wheat on a chessboard
// given that the number on each square doubles.
package grains

import (
	"errors"
)

const squares = 64

// Square returns how many grains were on each square.
func Square(square int) (uint64, error) {
	if square < 1 || square > squares {
		return 0, errors.New("square must be 1-64")
	}
	return 1 << uint(square-1), nil
}

// Total returns the total number of grains.
func Total() uint64 {
	return 1<<squares - 1
}
