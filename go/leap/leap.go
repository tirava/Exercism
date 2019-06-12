// leap implements calc of the leap year
package leap

// IsLeapYear returns true if the year is leap
func IsLeapYear(year int) bool {
	if year % 4 == 0 {
		if year % 100 !=0 || year % 400 == 0 {
			return true
		}
	}
	return false
}
