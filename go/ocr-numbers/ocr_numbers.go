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

func trimDigit(s string) string {
	s = strings.Trim(s, "\n")
	s = strings.TrimRight(s, " ")
	s = strings.Trim(s, "\n")
	return s
}

func recognizeDigit(s string) int {

	if v, ok := templates[trimDigit(s)]; ok {
		return v
	}

	//fmt.Println([]byte(trimDigit(s)))
	return -1
}

// Recognize returns recognized strings.
func Recognize(digits string) []string {

	result := make([]string, 0)
	numDigits := len(digits) / 12
	//fmt.Println(len(digits), numDigits)

	digit := ""
	for i := 0; i < numDigits; i++ { //1,5,9     1,11,18, 5,12,19
		digit = digits[i*numDigits+1 : i*numDigits+5]
		digit += digits[i*numDigits*3+5 : i*numDigits*3+5+4]
		digit += digits[i*numDigits*3*2+5+4 : i*numDigits*3*2+5+4+4]

		fmt.Println(digit)
	}

	d := recognizeDigit(digits)
	if d == -1 {
		result = append(result, "?")
	} else {
		result = append(result, strconv.Itoa(d))
	}

	return result
}
