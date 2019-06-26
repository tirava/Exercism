// Package diffsquares implements calculating difference between the square of the sum and the sum of the squares.
package diffsquares

// SquareOfSum returns Square Of Sum for the first N natural numbers.
func SquareOfSum(n int) int {
	sum := (n*n + n) / 2
	return sum * sum
}

// SumOfSquares returns Sum Of Squares for the first N natural numbers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns Difference between Square Of Sum and Sum Of Squares for the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
