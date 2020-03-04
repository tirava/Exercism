// Package minesweeper implements a minesweeper board.
package minesweeper

import (
	"errors"
	"fmt"
)

// Count counts mines around.
func (b Board) Count() error {
	lenX, lenY := len(b[0]), len(b)

	if lenX < 3 || lenY < 3 {
		return nil
	}

	for x := 1; x < lenX-1; x++ {
		for y := 1; y < lenY-1; y++ {
			s := b[y][x]

			if len(b[y]) != lenX {
				return errors.New("bad board line")
			}

			if s == ' ' {
				var count byte

				for xx := x - 1; xx <= x+1; xx++ {
					for yy := y - 1; yy <= y+1; yy++ {
						ss := b[yy][xx]

						if (xx == lenX || yy == lenY || xx == 0 || yy == 0) && (ss != '+' && ss != '-' && ss != '|') {
							return errors.New("bad board line")
						}

						if ss != '*' && ss != '+' && ss != '-' && ss != '|' && ss != ' ' && (ss < '0' || ss > '9') {
							fmt.Println("bad:", ss, string(ss), xx, yy)
							return errors.New("bad board")
						}

						if ss == '*' {
							count++
						}
					}
				}

				if count > 0 {
					b[y][x] = count + 48
				}
			}
		}
	}

	return nil
}
