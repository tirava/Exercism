// Package ocr implement digits recognition.
package ocr

import (
	"strings"
)

var templates = map[string]string{
	` _ 
| |
|_|`: "0",
	`   
  |
  |`: "1",
	` _ 
 _|
|_ `: "2",
	` _ 
 _|
 _|`: "3",
	`   
|_|
  |`: "4",
	` _ 
|_ 
 _|`: "5",
	` _ 
|_ 
|_|`: "6",
	` _ 
  |
  |`: "7",
	` _ 
|_|
|_|`: "8",
	` _ 
|_|
 _|`: "9",
}

func recognizeDigit(s string) string {

	if v, ok := templates[s]; ok {
		return v
	}

	return "?"
}

func recognizeDigits(digits string, numDigits int) string {
	sum := ""

	for i := 0; i < numDigits; i++ {
		digit := ""
		for j := 0; j < 3; j++ {
			start := (numDigits*3+1)*j + 3*i
			digit += digits[start:start+3] + "\n"
		}

		digit = strings.TrimRight(digit, "\n")

		d := recognizeDigit(digit)
		sum += d
	}

	return sum
}

// Recognize returns recognized strings.
func Recognize(digits string) []string {

	result := make([]string, 0)
	numDigits := len(digits) / 12

	digits = strings.TrimLeft(digits, "\n")

	for i := 1; i <= numDigits; i++ {
		crop := "\n" + strings.Repeat(" ", i*3) + "\n"
		if strings.Contains(digits, crop) {
			crop = strings.TrimRight(crop, "\n")
			slice := strings.Split(digits, crop)
			for i := 0; i < len(slice)-1; i++ {
				nd := len(slice[i])/12 + 1
				//slice[i] = strings.TrimRight(slice[i], "\n")
				slice[i] = strings.TrimLeft(slice[i], "\n")
				result = append(result, recognizeDigits(slice[i], nd))
				//fmt.Println([]byte(slice[i]))
				//fmt.Println(slice[i])
			}
			return result
		}
	}

	result = append(result, recognizeDigits(digits, numDigits))

	return result
}
