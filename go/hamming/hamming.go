// Package hamming implements calculating the Hamming difference between two DNA strands.
package hamming

import (
	"errors"
)

// Distance calculates Hamming distance or returns error.
func Distance(a, b string) (int, error) {

	lenA := len(a)
	if lenA != len(b) {
		return 0, errors.New("the strands are different length")
	}

	var num int
	for i := 0; i < lenA; i++ {
		if a[i] != b[i] {
			num += 1
		}
	}
	return num, nil
}
