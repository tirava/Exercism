// Package pythagorean implements finding Pythagorean triplets.
package pythagorean

type Triplet [3]int

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(sum int) (out []Triplet) {
	min, max := 1, sum
	out0 := append(out, Triplet{0, 0, 0})
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if c*c == a*a+b*b {
					if a+b+c == sum {
						out0 = append(out0, Triplet{a, b, c})
					}
				}
			}
		}
	}
	out = out0[1:]
	return
}

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (out []Triplet) {
	out0 := append(out, Triplet{0, 0, 0})
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if c*c == a*a+b*b {
					out0 = append(out0, Triplet{a, b, c})
				}
			}
		}
	}
	out = out0[1:]
	return
}
