// Package allyourbase implements converting a number,
// represented as a sequence of digits in one base, to any other base.
package allyourbase

// ConvertToBase returns 'a' converted to base 'b'.
func ConvertToBase(inBase int, inDigits []int, outBase int) (out []int, err error) {

	l := len(inDigits) - 1
	res := 0

	for i, d := range inDigits {
		//out = append(out, int(float64(d) * math.Pow(float64(inBase), float64(l - i))))
		res += d * intPow(inBase, l-i)
	}

	x1 := res
	//y := 0
	for i := 0; ; i++ {
		x2 := x1 % outBase
		//y += x2 * intPow(10, i) // -> res
		out = append(out, x2)
		x1 = x1 / outBase
		if x1 < outBase {
			//out = append(out, x2)
			break
		}
	}

	//out = append(out, res)

	return
}

func intPow(x, y int) int {
	if y == 0 {
		return 1
	}
	for ; y > 1; y-- {
		x *= x
	}
	return x
}
