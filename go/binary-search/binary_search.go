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
func SearchInts(data []int, x int) int {
	result := -1
	len := len(data) - 1

	for {
		len /= 2

		if len == 0 {
			if data[0] !=x {
				break
			} else {
				result = 0
			}
		}

		if
	}

	return result
}