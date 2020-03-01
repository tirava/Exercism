// Package binarysearch implements simple binary search.
package binarysearch

// SearchInts returns found position.
//func SearchInts(data []int, x int) int {
//	result := -1
//
//	for i, v := range data {
//		if v == x {
//			result = i
//		}
//	}
//
//	return result
//}
func SearchInts(list []int, x int) int {
	var low int
	high := len(list) - 1

	for {
		if low > high {
			break
		}

		mid := (low + high) / 2
		guess := list[mid]

		if guess == x {
			return mid
		}

		if guess > x {
			high = mid - 1
			continue
		}

		low = mid + 1
	}

	return -1
}
