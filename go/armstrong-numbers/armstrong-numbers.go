// Package armstrong implements answering is number Armstrong.
package armstrong

// IsNumber returns true if number is Armstrong.
func IsNumber(n int) bool {

	numbers := []int{}
	sum := 0

	for i := n; i > 0; i /= 10 {
		numbers = append(numbers, i%10)
	}

	for _, d := range numbers {
		pow := 1
		for i := 0; i < len(numbers); i++ {
			pow *= d
		}
		sum += pow
	}

	return sum == n
}
