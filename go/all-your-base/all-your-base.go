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

	res := 0

	for _, d := range inDigits {
		if d < 0 || d >= inBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		res = d + res*inBase
	}

	for i := 0; ; i++ {
		rem := res % outBase
		out = append([]int{rem}, out...)
		res = res / outBase
		if res == 0 {
			break
		}
	}

	return
}
