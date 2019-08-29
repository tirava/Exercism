// Package palindrome implements detecting palindrome products in a given range.
package palindrome

import (
	"errors"
	"math"
	"strconv"
)

// Product is the base type for products.
type Product struct {
	Product        int
	Factorizations [][2]int
}

type palindrome struct {
	x, y []int
	p    bool
}

// Products return min & max products.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {

	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	products := make(map[int]palindrome)

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			k := i * j
			products[k] = palindrome{append(products[k].x, i), append(products[k].y, j), isPalindrom(strconv.Itoa(k))}
		}
	}

	min, max := math.MaxInt32, -math.MaxInt32
	for i, v := range products {
		if v.p {
			if min > i {
				min = i
			}
			if max < i {
				max = i
			}
		}
	}

	if max == -math.MaxInt32 {
		return Product{}, Product{}, errors.New("no palindromes")
	}

	pmin.Product, pmax.Product = min, max

	for i := 0; i < len(products[min].x); i++ {
		pmin.Factorizations = append(pmin.Factorizations, [2]int{products[min].x[i], products[min].y[i]})
	}
	for i := 0; i < len(products[max].x); i++ {
		pmax.Factorizations = append(pmax.Factorizations, [2]int{products[max].x[i], products[max].y[i]})
	}

	return
}

func isPalindrom(s string) bool {

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return s == string(r)
}
