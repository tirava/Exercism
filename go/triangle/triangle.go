// Package triangle determine if a triangle is equilateral, isosceles, or scalene.
package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides checks type of triangle.
func KindFromSides(a, b, c float64) (k Kind) {

	if a <= 0 || b <= 0 || c <= 0 {
		return // NaT
	}

	if (a+b) < c || (b+c) < a || (a+c) < b {
		return // NaT
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return // NaT
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return // NaT
	}

	if a == b && b == c && a == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else {
		k = Sca
	}

	return
}
