// Package darts implements darts playing.
package darts

import (
	"math"
)

// Score returns points for Decart x,y coordinates in circle.
func Score(x, y float64) int {

	point := math.Sqrt(x*x + y*y)
	//point := math.Hypot(x, y)

	switch {
	case point > 10:
		return 0
	case point <= 10 && point > 5:
		return 1
	case point <= 5 && point > 1:
		return 5
	case point <= 1:
		return 10
	default:
		return 0
	}
}
