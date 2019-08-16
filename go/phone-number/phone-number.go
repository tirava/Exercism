// Package phonenumber implements cleaning up user-entered phone numbers
// so that they can be sent SMS messages.
package phonenumber

import (
	"errors"
	"regexp"
	"unicode"
)

// Number returns clean number.
func Number(n string) (string, error) {

	nums := ""
	for _, s := range n {
		if unicode.IsDigit(s) {
			nums += string(s)
		}
	}
	if len(nums) == 11 && nums[0] == '1' {
		nums = nums[1:]
	}

	re := regexp.MustCompile(`^2\d{2}[2-9]\d{6}$`)
	if !re.MatchString(nums) {
		return "", errors.New("not phone number")
	}

	return nums, nil
}

// Format returns formatted number.
func Format(n string) (string, error) {
	s, err := Number(n)
	if err != nil {
		return "", err
	}
	return "(" + s[:3] + ") " + s[3:6] + "-" + s[6:], nil
}

// AreaCode returns clean numbers.
func AreaCode(n string) (string, error) {
	s, err := Number(n)
	if err != nil {
		return "", err
	}
	return s[:3], nil
}
