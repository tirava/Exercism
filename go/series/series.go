// Package series implements output all the contiguous substrings
// of length n in that string in the order that they appear.
package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) (res []string) {
	for i := 0; i <= len(s)-n; i++ {
		res = append(res, UnsafeFirst(n, s[i:]))
	}
	return

}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First is safe version of the UnsafeFirst.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return UnsafeFirst(n, s), true
}
