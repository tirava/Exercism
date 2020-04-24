// Package alphametics implements puzzle where letters in words are replaced with numbers.
package alphametics

import (
	"fmt"
	"strings"
)

// Solve solves alphametics puzzles.
func Solve(puzzle string) (map[string]int, error) {
	sep := strings.Split(puzzle, "==")
	ops := strings.Split(sep[0], "+")

	op1 := strings.TrimSpace(ops[0])
	op2 := strings.TrimSpace(ops[1])
	res := strings.TrimSpace(sep[1])

	fmt.Println(sep, ops, op1, op2, res)

	return nil, nil
}
