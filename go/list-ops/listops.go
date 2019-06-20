// Package listops implement basic list operations.
package listops

// IntList is a common type for lists.
type IntList []int

// binFunc implements binary functions for lists.
type binFunc func(int, int) int

// unaryFunc implements unary function for lists
type unaryFunc func(int) int

// predFunc implements predicated functions for lists
type predFunc func(int) bool

// Length returns list length.
func (il IntList) Length() int {
	return len(il)
}

// Foldr returns list fold right.
func (il IntList) Foldr(bf binFunc, in int) int {
	for i := range il {
		in = bf(il[len(il)-1-i], in)
	}
	return in
}

// Foldl returns list fold left.
func (il IntList) Foldl(bf binFunc, in int) int {
	for _, x := range il {
		in = bf(in, x)
	}
	return in
}

// Reverse returns reverse list.
func (il IntList) Reverse() IntList {
	reversed := IntList{}
	for i := range il {
		n := il[len(il)-1-i]
		reversed = append(reversed, n)
	}
	return reversed
}

// Append returns appended lists.
func (il IntList) Append(in []int) IntList {
	return append(il, in...)
}

// Concat returns concatenated lists.
func (il IntList) Concat(args []IntList) IntList {
	for _, in := range args {
		il = append(il, in...)
	}
	return il
}

// Map returns map of list after unary function.
func (il IntList) Map(uf unaryFunc) IntList {
	for i, n := range il {
		il[i] = uf(n)
	}
	return il
}

// Filter returns map of list after predicated function.
func (il IntList) Filter(pf predFunc) IntList {
	filtered := IntList{}
	for _, n := range il {
		if pf(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}
