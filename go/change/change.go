// Package change implements coins to be given to a customer.
package change

import "fmt"

var result = make([][]int, 0)

// Change returns coins.
func Change(cost []int, coin int) ([]int, error) {
	var indexCount int
	minCount := 99999
	N := len(cost)
	count := make([]int, N)

	change(N-1, coin, count, cost)

	for i, counts := range result {
		var min int
		for _, count := range counts {
			min += count
		}

		if min < minCount {
			minCount = min
			indexCount = i
		}
	}

	fmt.Println(result[indexCount])

	return count, nil
}

func change(maxcoin, sum int, count, cost []int) {

	if sum == 0 {
		//fmt.Println(maxcoin, sum, count, cost)
		copyCount := make([]int, len(count))
		copy(copyCount, count)
		result = append(result, copyCount)
		return
	}

	if sum >= cost[maxcoin] {
		count[maxcoin]++
		change(maxcoin, sum-cost[maxcoin], count, cost)
		count[maxcoin]--
	}

	if maxcoin != 0 {
		change(maxcoin-1, sum, count, cost)
	}
}
