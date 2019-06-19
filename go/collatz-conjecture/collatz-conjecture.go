// Package collatzconjecture implements Collatz Conjecture.
package collatzconjecture

import "errors"

// Collatz Conjecture calculates steps to get "1".
func CollatzConjecture(input int) (steps int, err error) {
	if input <= 0 {
		return -1, errors.New("number must be more zero")
	}

	if input == 1 {
		return 0, nil
	}

	for {
		if input%2 == 0 {
			input = input / 2
		} else {
			input = input*3 + 1
		}
		steps++
		if input == 1 {
			break
		}
	}

	return steps, nil
}
