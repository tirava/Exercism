// Package roman-numerals implements conversion from normal numbers to Roman Numerals.
package romannumerals

import "errors"

// ToRomanNumeral converts Arabic to Roman Numerals.
func ToRomanNumeral(a int) (string, error) {
	if a <= 0 || a > 3000 {
		return "", errors.New("there is no need to be able to convert numbers larger than about 3000")
	}

	r := hirelo(a/1000, "", "", "M")
	a %= 1000
	r += hirelo(a/100, "M", "D", "C")
	a %= 100
	r += hirelo(a/10, "C", "L", "X")
	a %= 10
	r += hirelo(a, "X", "V", "I")

	return r, nil
}

func hirelo(num int, hi, re, lo string) string {
	switch num {
	case 9:
		return lo + hi
	case 8:
		return re + lo + lo + lo
	case 7:
		return re + lo + lo
	case 6:
		return re + lo
	case 5:
		return re
	case 4:
		return lo + re
	case 3:
		return lo + lo + lo
	case 2:
		return lo + lo
	case 1:
		return lo
	}
	return ""
}
