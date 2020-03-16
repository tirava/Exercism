// Package change implements coins to be given to a customer.
package change

import "errors"

//var result = make([][]int, 0)

// Change returns coins.
func Change(coins []int, amount int) ([]int, error) {
	if amount < 0 {
		return nil, errors.New("amount must be >= 0")
	}

	amounts := make([][]int, amount+1)
	amounts[0] = []int{}

	for i := range amounts {
		for _, coin := range coins {
			if i-coin >= 0 && amounts[i-coin] != nil && (amounts[i] == nil || len(amounts[i-coin])+1 < len(amounts[i])) {
				amounts[i] = append([]int{coin}, amounts[i-coin]...)
			}
		}
	}

	if amounts[amount] == nil {
		return nil, errors.New("can not change")
	}

	return amounts[amount], nil
}

//func Change(cost []int, coin int) ([]int, error) {
//	var indexCount int
//	minCount := 99999
//	N := len(cost)
//	count := make([]int, N)
//	out := make([]int, 0)
//
//	result = result[:0]
//	change(N-1, coin, count, cost)
//
//	for i, counts := range result {
//		var min int
//		for _, count := range counts {
//			min += count
//		}
//
//		if min < minCount {
//			minCount = min
//			indexCount = i
//		}
//	}
//
//	//fmt.Println(result[indexCount])
//	for k, v := range result[indexCount] {
//		for i := 0; i < v; i++ {
//			out = append(out, cost[k])
//		}
//	}
//
//	return out, nil
//}
//
//func change(maxcoin, sum int, count, cost []int) {
//
//	if sum == 0 {
//		//fmt.Println(maxcoin, sum, count, cost)
//		copyCount := make([]int, len(count))
//		copy(copyCount, count)
//		result = append(result, copyCount)
//		return
//	}
//
//	if sum >= cost[maxcoin] {
//		count[maxcoin]++
//		change(maxcoin, sum-cost[maxcoin], count, cost)
//		count[maxcoin]--
//	}
//
//	if maxcoin != 0 {
//		change(maxcoin-1, sum, count, cost)
//	}
//}
