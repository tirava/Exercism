// Package change implements coins to be given to a customer.
package change

import "fmt"

// Change returns coins.
func Change(cost []int, coin int) ([]int, error) {
	count := make([]int, len(cost))
	N := len(cost)

	change(N-1, coin, count, cost)

	return count, nil
}

func change(maxcoin, sum int, count, cost []int) {

	if sum == 0 {
		fmt.Println(maxcoin, sum, count, cost)
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
