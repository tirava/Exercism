// Package prime implements nth Prime.
package prime

const maxPrimeIndex = 105000 + 1 // max prime

// Nth returns nth Prime or false.
func Nth(n int) (p int, ok bool) {

	if n == 0 {
		return 0, false
	}

	boolArr := make([]bool, maxPrimeIndex) //max prime index for first 500 numbers
	boolArr[0] = true                      // non prime
	boolArr[1] = true                      // non prime

	// fill bool array non prime numbers, prime = false for default
	for i := 2; i < maxPrimeIndex; i++ {
		if !boolArr[i] {
			if i*i < maxPrimeIndex {
				for j := i * i; j < maxPrimeIndex; j += i {
					boolArr[j] = true
				}
			}
		}
	}

	// convert bool indexes to prime numbers
	primeArr := make([]int, 0, n)
	for i := 0; i < maxPrimeIndex; i++ {
		if !boolArr[i] {
			primeArr = append(primeArr, i)
			n--
			if n == 0 {
				break // no need more than n
			}
		}
	}

	return primeArr[len(primeArr)-1], true
}
