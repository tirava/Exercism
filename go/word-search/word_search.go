// Package wordsearch implements word search puzzles you get a square of letters
// and have to find specific words in them.
package wordsearch

import (
	"errors"
	"fmt"
	"strings"
)

// Solve returns words coordinates.
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)

	search1 := search(words, puzzle, result, false, false)
	var search2, search3 bool
	if len(words) > 1 {
		puzzleReverse := reversePuzzle(puzzle)
		search2 = search(words, puzzleReverse, result, true, false)
		puzzleDiagonal := diagonalPuzzle(puzzle)
		search3 = search(words, puzzleDiagonal, result, false, true)
	}

	if !search1 && !search2 && !search3 {
		return nil, errors.New("not found")
	}

	return result, nil
}

func diagonalPuzzle(words []string) []string {
	newWords := make([]string, len(words[0])) //+len(words)-3)
	newWord := strings.Builder{}

	for i1, j1, i2, j2 := len(words[0])-2, 0, len(words[0])-1, 0; i1 >= 0; i1, j2 = i1-1, j2+1 {
		for i, j := i1, j1; i <= i2; i, j = i+1, j+1 {
			newWord.WriteByte(words[j][i])
		}

		newWords[j2] = newWord.String()
		fmt.Println(newWord.String())
		newWord.Reset()
	}

	return newWords
}

func reversePuzzle(words []string) []string {
	newWords := make([]string, len(words[0]))
	newWord := strings.Builder{}

	for i := 0; i < len(words[0]); i++ {
		for j := 0; j < len(words); j++ {
			newWord.WriteByte(words[j][i])
		}

		newWords[i] = newWord.String()
		newWord.Reset()
	}

	return newWords
}

func search(words []string, puzzle []string, result map[string][2][2]int, reversed, diag bool) bool {
	var count int

	for i, pz := range puzzle {
		for _, w := range words {
			begin := strings.Index(pz, w)
			if begin != -1 {
				var deltaDiag1, deltaDiag2 int
				if diag {
					deltaDiag1 = begin //len(pz) - 2 // - i
					deltaDiag2 = deltaDiag1 + len(w) - 1
					result[w] = [2][2]int{{begin, begin}, {begin + len(w) - 1, begin + len(w) - 1}}
					fmt.Print("diags:", deltaDiag1, deltaDiag2)
				} else {
					if !reversed {
						result[w] = [2][2]int{{begin, i}, {begin + len(w) - 1, i}}
					} else {
						result[w] = [2][2]int{{i, begin}, {i, begin + len(w) - 1}}
					}
				}
				count++
			}

			begin = strings.Index(pz, reverse(w))
			if begin != -1 {
				if !reversed {
					result[w] = [2][2]int{{begin + len(w) - 1, i}, {begin, i}}
				} else {
					result[w] = [2][2]int{{i, begin + len(w) - 1}, {i, begin}}
				}
				count++
			}
		}
	}

	return count != 0
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//package wordsearch
//
//import (
//"errors"
//)
//
///**
//1 1 1
//1 2 1
//1 1 1
//
//2 is the center
//**/
//const (
//	C1 = iota // (0, 0)
//	C2        // (0, 1)
//	C3        // (0, 2)
//	C4        // (1, 2)
//	C5        // (2, 2)
//	C6        // (2, 1)
//	C7        // (2, 0)
//	C8        // (1, 0)
//)
//
//func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
//
//	res := make(map[string][2][2]int)
//
//	bts := make([][]byte, len(puzzle))
//
//	for i := 0; i < len(bts); i++ {
//		bts[i] = []byte(puzzle[i])
//	}
//
//	for _, v := range words {
//		t, ok := findPos(bts, []byte(v))
//		if ok {
//			res[v] = t
//		} else {
//			return res, errors.New("fail to locate a word that is not in the puzzle")
//		}
//	}
//
//	return res, nil
//
//}
//
//func findPos(bts [][]byte, dest []byte) ([2][2]int, bool) {
//	for i := 0; i < len(bts); i++ {
//		for j := 0; j < len(bts[0]); j++ {
//			if bts[i][j] == dest[0] {
//				for k := C1; k <= C8; k++ {
//					if isExist(bts, dest, 0, i, j, k) {
//						endR, endC := caclPos(i, j, k, len(dest)-1)
//						return [2][2]int{{j, i}, {endC, endR}}, true
//					}
//				}
//
//			}
//		}
//	}
//
//	return [2][2]int{}, false
//}
//
//func caclPos(row, col, c, dx int) (destR, destC int) {
//	switch c {
//	case C1:
//		destR, destC = row-dx, col-dx
//	case C2:
//		destR, destC = row-dx, col
//	case C3:
//		destR, destC = row-dx, col+dx
//	case C4:
//		destR, destC = row, col+dx
//	case C5:
//		destR, destC = row+dx, col+dx
//	case C6:
//		destR, destC = row+dx, col
//	case C7:
//		destR, destC = row+dx, col-dx
//	case C8:
//		destR, destC = row, col-dx
//	}
//	return
//}
//
//func isExist(bts [][]byte, dest []byte, index, row, col, c int) bool {
//
//	if index == len(dest) {
//		return true
//	}
//
//	if row < 0 || row >= len(bts) || col < 0 || col >= len(bts[0]) {
//		return false
//	}
//
//	if bts[row][col] != dest[index] {
//		return false
//	}
//
//	destR, destC := caclPos(row, col, c, 1)
//
//	return isExist(bts, dest, index+1, destR, destC, c)
//}
