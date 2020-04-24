// Package alphametics implements puzzle where letters in words are replaced with numbers.
package alphametics

import (
	"errors"
	"strings"
)

var f10 = fact(10)

type result struct {
	hash  map[string]int
	zero  bool
	found bool
}

// Solve solves alphametics puzzles.
func Solve(puzzle string) (map[string]int, error) {
	hash := make(map[string]int)
	slice := make([]string, 0, 10)
	//f10 := fact(10)

	sep := strings.Split(puzzle, "==")
	ops := strings.Split(sep[0], "+")

	if len(ops) < 2 {
		return nil, errors.New("need 2+ operands")
	}

	res := strings.TrimSpace(sep[1])

	for i, op := range ops {
		ops[i] = strings.TrimSpace(op)
	}

	allString := strings.Join(ops, "") + res

	for i, s := range allString {
		if _, ok := hash[string(s)]; !ok {
			hash[string(s)] = i
			slice = append(slice, string(s))
		}
	}

	for len(slice) < 10 {
		slice = append(slice, "_")
	}

	result := worker(hash, slice, ops, res)

	if result.zero {
		return nil, errors.New("no leading zero")
	}

	return result.hash, nil
}

func worker(hash map[string]int, slice, ops []string, res string) result {
	ch := make(chan result, 8)

	//result1 := job(0, f10/2, hash, slice, ops, res)
	//result2 := job(f10/2, f10, hash, slice, ops, res)
	go job(0, f10/8, hash, slice, ops, res, ch, 1)
	go job(f10/8+1, f10*2/8, hash, slice, ops, res, ch, 2)
	go job(f10*2/8+1, f10*3/8, hash, slice, ops, res, ch, 3)
	go job(f10*3/8+1, f10*4/8, hash, slice, ops, res, ch, 4)
	go job(f10*4/8+1, f10*5/8, hash, slice, ops, res, ch, 5)
	go job(f10*5/8+1, f10*6/8, hash, slice, ops, res, ch, 6)
	go job(f10*6/8+1, f10*7/8, hash, slice, ops, res, ch, 7)
	go job(f10*7/8+1, f10, hash, slice, ops, res, ch, 8)

	result := result{}

	for r := range ch {
		if r.found {
			result.hash = r.hash
			return result
		}
		if r.zero {
			result.zero = true
			return result
		}
	}

	return result
}

func job(index1, index2 int, hash map[string]int, slice, ops []string, res string, ch chan result, num int) {
	var result result

	result.hash = make(map[string]int, len(hash))

	for k, v := range hash {
		result.hash[k] = v
	}

	//fmt.Println("job:", num, "from:", index1, "to:", index2)

LOOP:
	for i := index1; i < index2; i++ {
		permSlice := permStr(i, slice)

		var sumOps int
		for _, op := range ops {
			sumOps += getNumber(permSlice, op)
		}

		sumRes := getNumber(permSlice, res)

		if sumOps == sumRes {

			for i, s := range permSlice {
				if s == "_" {
					continue
				}

				result.hash[s] = i
			}

			if result.hash[res[0:1]] == 0 {
				result.zero = true
				continue
			}

			for _, op := range ops {
				if result.hash[op[0:1]] == 0 {
					result.zero = true
					continue LOOP
				}
			}

			result.zero = false
			result.found = true

			//fmt.Println("job:", num, i, permSlice, "sumRes:", sumRes, "sumOps:", sumOps)
			break LOOP
		}

		//fmt.Printf("%d: %d of %d\r", num, i, f10)
	}

	ch <- result
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
