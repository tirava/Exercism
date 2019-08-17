// Package prime implements computing the prime factors of a given natural number.
package prime

import "math"

// Factors returns slice of factors.
func Factors(n int64) (res []int64) {

	res = make([]int64, 8)
	var i int64

	for i = 2; i < math.MaxInt64; i++ {
		for {
			if n%i == 0 {
				n = n / i
				res = append(res, i)
			} else {
				break
			}
		}
	}

	return
}
