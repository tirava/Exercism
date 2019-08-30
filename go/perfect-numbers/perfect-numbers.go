// Package perfect implements determining if a number is perfect, abundant, or deficient.
package perfect

import "errors"

// Classification is the base type for numbers
type Classification string

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

	return ClassificationPerfect, nil
}
