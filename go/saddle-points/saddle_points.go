package matrix

import "fmt"

// Pair is the base pair type.
type Pair [2]int

func (m Matrix) Saddle() []Pair {
	//maxRows := make([]Pair, 0)
	cols := m.Cols()

	X, Y := -1, -1
	for _, row := range m.Rows() {
		maxVal := -999999
		for x, val := range row {
			if val >= maxVal {
				maxVal = val
				X = x
				minVal := 999999
				for y, val := range cols[x] {
					if val <= minVal {
						minVal = val
						Y = y
					}
				}
			}
		}
		fmt.Println(X, Y)
	}

	//for _, col := range m.Cols() {
	//	minVal := 999999
	//	for y, val := range col {
	//		if val <= minVal {
	//			minVal = val
	//			Y = y
	//		}
	//	}
	//	fmt.Println(Y)
	//}

	//fmt.Println(X, Y)
	fmt.Println(m.Rows())
	fmt.Println(m.Cols())
	return nil
}
