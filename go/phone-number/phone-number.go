// Package phonenumber implements cleaning up user-entered phone numbers
// so that they can be sent SMS messages.
package phonenumber

import (
	"errors"
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

	//re, err := regexp.Compile(`^2$`)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////println(nums)
	//
	//if !re.MatchString(n) {
	//	return "", errors.New("not phone number")
	//}

	if len(nums) < 10 || len(nums) > 11 {
		return "", errors.New("not phone number")
	}

	if len(nums) == 11 && nums[0] != '1' {
		return "", errors.New("not phone number")
	}

	if len(nums) == 10 && nums[0] != '2' {
		return "", errors.New("not phone number")
	}

	if len(nums) == 10 && nums[3] == '0' {
		return "", errors.New("not phone number")
	}

	if len(nums) == 10 && nums[3] == '1' {
		return "", errors.New("not phone number")
	}

	if len(nums) == 11 && nums[0] == '1' {
		return nums[1:], nil
	}

	return nums, nil
}

// Number returns formatted number.
func Format(n string) (string, error) {

	return "", nil
}

// Number returns clean numbers.
func AreaCode(n string) (string, error) {

	return "", nil
}
