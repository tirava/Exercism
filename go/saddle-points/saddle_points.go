package matrix

// Pair is the base pair type.
type Pair [2]int

// Saddle returns saddle points.
func (m Matrix) Saddle() []Pair {
	saddlePoints := make([]Pair, 0)
	maxRow := make([]Pair, 0)
	minCol := make([]Pair, 0)

	for Y, row := range m.Rows() {
		maxVal := -999999
		for _, valX := range row {
			if valX > maxVal {
				maxVal = valX
			}
		}
		for X, valX := range row {
			if valX == maxVal {
				maxRow = append(maxRow, Pair{Y, X})
			}
		}
	}

	for X, col := range m.Cols() {
		minVal := 999999
		for _, valY := range col {
			if valY < minVal {
				minVal = valY
			}
		}
		for Y, valY := range col {
			if valY == minVal {
				minCol = append(minCol, Pair{Y, X})
			}
		}
	}

	for _, pairRow := range maxRow {
		for _, pairCol := range minCol {
			if pairRow == pairCol {
				saddlePoints = append(saddlePoints, pairRow)
			}
		}
	}

	return saddlePoints
}
