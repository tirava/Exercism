// Package minesweeper implements a minesweeper board.
package minesweeper

// Count counts mines around.
func (b Board) Count() error {
	lenX, lenY := len(b[0]), len(b)
	if lenX < 3 || lenY < 3 {
		return nil
	}

	for x := 1; x < lenX-1; x++ {
		for y := 1; y < lenY-1; y++ {
			s := b[y][x]

			if s == ' ' {
				var count byte

				for xx := x - 1; xx <= x+1; xx++ {
					for yy := y - 1; yy <= y+1; yy++ {
						ss := b[yy][xx]

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
