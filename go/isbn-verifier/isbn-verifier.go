// Package isbn implements checking if the provided string is a valid ISBN-10.
package isbn

import "strings"

// IsValidISBN checks if isbn is true.
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) == 0 || len(isbn) != 10 {
		return false
	}
	var sum int
	for i, s := range isbn {
		d := s - 48
		if s == 'X' {
			if i != len(isbn)-1 {
				return false
			}
			d = 10
		}
		if d < 0 || d > 9 && s != 'X' {
			return false
		}
		sum += (10 - i) * int(d)
	}
	if sum%11 == 0 {
		return true
	}
	return false
}
