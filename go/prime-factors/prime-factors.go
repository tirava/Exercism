// Package prime implements computing the prime factors of a given natural number.
package prime

// Factors returns slice of factors.
func Factors(n int64) []int64 {

	res := []int64{}
	var i int64

	for i = 2; i <= n; {
		if n%i == 0 {
			n = n / i
			res = append(res, i)
			continue
		}
		i++
	}

	return res
}
