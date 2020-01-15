// Package ocr implement digits recognition.
package ocr

import (
	"fmt"
	"strconv"
	"strings"
)

var templates = map[string]int{
	` _ 
| |
|_|`: 0,
	`   
  |
  |`: 1,
	` _ 
 _|
|_ `: 2,
	` _ 
 _|
 _|`: 3,
	`   
|_|
  |`: 4,
	` _ 
|_ 
 _|`: 5,
	` _ 
|_ 
|_|`: 6,
	` _ 
  |
  |`: 7,
	` _ 
|_|
|_|`: 8,
	` _ 
|_|
 _|`: 9,
}

func recognizeDigit(s string) int {

	if v, ok := templates[s]; ok {
		return v
	}

	return -1
}

// Recognize returns recognized strings.
func Recognize(digits string) []string {

	result := make([]string, 0)
	numDigits := len(digits) / 12

	digits = strings.TrimLeft(digits, "\n")
	for i := 0; i < numDigits; i++ {
		digit := ""
		for j := 0; j < 3; j++ {
			start := (numDigits*3+1)*j + 3*i
			digit += digits[start:start+3] + "\n"
		}

		digit = strings.TrimRight(digit, "\n")
		fmt.Println(digit)

		d := recognizeDigit(digit)
		if d == -1 {
			result = []string{"?"}
			return result
		}
		result = append(result, strconv.Itoa(d))
	}

	return result
}
