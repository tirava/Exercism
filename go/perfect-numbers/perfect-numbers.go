// Package perfect implements determining if a number is perfect, abundant, or deficient.
package perfect

import (
	"errors"
)

// Classification is the base type for numbers.
type Classification string

// Classification constants.
const (
	ClassificationPerfect   = "Perfect"
	ClassificationDeficient = "Deficient"
	ClassificationAbundant  = "Abundant"
)

// ErrOnlyPositive is error for not positive numbers.
var ErrOnlyPositive = errors.New("not a positive number")

// Classify returns classified number.
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}

	var sum int64

	for i := n / 2; i > 1; i-- {
		if n%i == 0 {
			sum += i
		}
	}
	sum++

	switch {
	case sum-n > 0:
		return ClassificationAbundant, nil
	case sum-n == 0 && n != 1:
		return ClassificationPerfect, nil
	default:
		return ClassificationDeficient, nil
	}
}
