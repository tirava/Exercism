// Package ocr implement digits recognition.
package ocr

import (
	"strconv"
	"strings"
)

var templates = map[string]int{
	`
 _ 
| |
|_|
   `: 0,
	`
   
  |
  |
   `: 1,
	`
 _ 
 _|
|_ 
   `: 2,
	`
 _ 
 _|
 _|
   `: 3,
	`
   
|_|
  |
   `: 4,
	`
 _ 
|_ 
 _|
   `: 5,
	`
 _ 
|_ 
|_|
   `: 6,
	`
 _ 
  |
  |
   `: 7,
	`
 _ 
|_|
|_|
   `: 8,
	`
 _ 
|_|
 _|
   `: 9,
}

func trimDigit(s string) string {
	s = strings.Trim(s, "\n")
	s = strings.TrimRight(s, " ")
	s = strings.Trim(s, "\n")
	return s
}

func recognizeDigit(s string) int {

	if v, ok := templates[s]; ok {
		return v
	}

	//fmt.Println([]byte(trimDigit(s)))
	return -1
}

// Recognize returns recognized strings.
func Recognize(digits string) []string {

	result := make([]string, 0)

	d := recognizeDigit(digits)
	if d == -1 {
		result = append(result, "?")
	} else {
		result = append(result, strconv.Itoa(d))
	}

	return result
}
