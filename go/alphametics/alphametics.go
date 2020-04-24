// Package alphametics implements puzzle where letters in words are replaced with numbers.
package alphametics

import (
	"errors"
	"fmt"
	"strings"
)

// Solve solves alphametics puzzles.
func Solve(puzzle string) (map[string]int, error) {
	hash := make(map[string]int)
	slice := make([]string, 0, 10)
	f10 := fact(10)

	sep := strings.Split(puzzle, "==")
	ops := strings.Split(sep[0], "+")

	if len(ops) < 2 {
		return nil, errors.New("need 2+ operands")
	}

	op1 := strings.TrimSpace(ops[0])
	op2 := strings.TrimSpace(ops[1])
	res := strings.TrimSpace(sep[1])

	for i, s := range op1 + op2 + res {
		if _, ok := hash[string(s)]; !ok {
			hash[string(s)] = i
			slice = append(slice, string(s))
		}
	}

	for len(slice) < 10 {
		slice = append(slice, "_")
	}

	var permSlice []string
	for i := 0; i < f10; i++ {
		permSlice = permStr(i, slice)

		num1 := getNumber(permSlice, op1)
		num2 := getNumber(permSlice, op2)
		sum3 := getNumber(permSlice, res)
		sum := num1 + num2

		if sum3 == sum {
			fmt.Println(i, permSlice, "num1:", num1, "num2:", num2, "sum3:", sum3, "real sum:", sum)
			break
		}

		fmt.Printf("%d of %d\r", i, f10)
	}

	for i, s := range permSlice {
		if s == "_" {
			continue
		}

		hash[s] = i
	}

	return hash, nil
}

func getNumber(slice []string, str string) int {
	var res int

	for _, s := range str {
		for k, v := range slice {
			if v == string(s) {
				res = res*10 + k
			}
		}
	}

	return res
}

func permStr(index int, src []string) []string {
	res := make([]string, 0, len(src))
	source := make([]string, 0, len(src))

	for _, pos := range src {
		source = append(source, pos)
	}

	for j := 0; j < len(src); j++ {
		p := (index / fact(len(src)-1-j)) % len(source)
		res = append(res, source[p])
		source = append(source[:p], source[p+1:]...)
	}

	return res
}

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
