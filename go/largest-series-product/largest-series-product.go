// Package lsproduct implements calculating the largest product for a contiguous substring of digits.
package lsproduct

import "errors"

// LargestSeriesProduct calculates the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (product int64, err error) {

	var curProd, maxProd int64
	curProd = 1

	if span > len(digits) || span < 0 {
		return -1, errors.New("span must be greater than zero or span must be smaller than string length")
	}

	for i := 0; i <= len(digits)-span; i++ {
		curProd = 1
		for j := i; j < i+span; j++ {
			//println("i, j", i, j)
			n := digits[j] - 48
			if n < 0 || n > 9 {
				return -1, errors.New("digits input must only contain digits")
			}
			curProd *= int64(n)
			//println("curProd=", curProd)
		}
		if curProd > maxProd {
			maxProd = curProd
			curProd = 1
		}
	}

	return maxProd, nil
}
