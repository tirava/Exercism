// Package allyourbase implements converting a number,
// represented as a sequence of digits in one base, to any other base.
package allyourbase

import (
	"errors"
)

// ConvertToBase returns 'a' converted to base 'b'.
func ConvertToBase(inBase int, inDigits []int, outBase int) (out []int, err error) {

	if inBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	l := len(inDigits) - 1
	res := 0

	for i, d := range inDigits {
		if d < 0 || d >= inBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		res += d * intPow(inBase, l-i)
	}

	for i := 0; ; i++ {
		rem := res % outBase
		out = append(out, rem)
		res = res / outBase
		if res == 0 {
			break
		}
	}

	l = len(out)
	if l > 1 {
		for i := 0; i < l/2; i++ {
			out[i], out[l-i-1] = out[l-i-1], out[i]
		}
	}

	return
}

func intPow(x, y int) int {
	res := x
	if y == 0 {
		return 1
	}
	for ; y > 1; y-- {
		res *= x
	}
	return res
}
