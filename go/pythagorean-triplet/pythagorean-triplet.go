// Package pythagorean implements finding Pythagorean triplets.
package pythagorean

type Triplet [3]int

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(sum int) (out []Triplet) {
	for _, x := range Range(1, sum/2) {
		if x[0]+x[1]+x[2] == sum {
			out = append(out, x)
		}
	}
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
