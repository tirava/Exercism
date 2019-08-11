// Package sieve implements Sieve of Eratosthenes to find all the primes from 2 up to a given number.
package sieve

// Sieve fills slice of prime numbers not more than limit.
func Sieve(limit int) (result []int) {

	if limit < 2 {
		return
	}

	boolArr := make([]bool, limit+2)
	boolArr[0], boolArr[1] = true, true // non prime

	// fill bool array non prime numbers, prime = false for default
	for i := 2; i < limit+2; i++ {
		if !boolArr[i] {
			if i*i < limit {
				for j := i * i; j < limit+2; j += i {
					boolArr[j] = true
				}
			}
		}
	}

	// convert bool indexes to prime numbers
	for i := 0; i < limit+2; i++ {
		if !boolArr[i] {
			if i > limit {
				break // no need more than limit
			}
			result = append(result, i)
		}
	}

	return
}
