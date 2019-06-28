// Package summultiples implements finding the sum of all the unique multiples of particular numbers up to
// but not including that number.
package summultiples

// SumMultiples returns sum of all the unique multiples.
func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		for _, n := range divisors {
			if n != 0 && i%n == 0 {
				sum += i
				break
			}
		}
	}

	return sum
}
