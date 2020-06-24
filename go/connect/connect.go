// Package connect provide the function needed to verify the winner of the hex game.
package connect

var pos = [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 0}, {1, -1}, {0, -1}}

// MatrixVisit returns true if a path exists horizontally of the marker sign
func MatrixVisit(m []string, marker uint8) bool {
	lrm := len(m)
	lcm := len(m[0])
	// initialize visited
	visited := make([][]bool, lrm)
	for i := 0; i < lrm; i++ {
		visited[i] = make([]bool, lcm)
	}
	// bfs visit
	mem := make([][]int, 0, lrm)
	for i := 0; i < lrm; i++ {
		if m[i][0] == marker {
			mem = append(mem, []int{i, 0})
		}
	}
	for len(mem) > 0 {
		u := mem[0]
		mem = mem[1:]
		visited[u[0]][u[1]] = true
		if u[1] == lcm-1 {
			return true
		}
		// add nodes to visit
		for _, dir := range pos {
			newVisit := []int{u[0] + dir[0], u[1] + dir[1]}
			if newVisit[0] >= 0 && newVisit[0] < lrm && newVisit[1] >= 0 && newVisit[1] < lcm &&
				m[newVisit[0]][newVisit[1]] == marker && !visited[newVisit[0]][newVisit[1]] {
				mem = append(mem, newVisit)
			}
		}
	}
	return false
}

func matrixRotate(m []string) (res []string) {
	lrm := len(m)
	lcm := len(m[0])
	// initialize
	res = make([]string, 0, lcm)
	for j := 0; j < lcm; j++ {
		tmp := make([]byte, lrm)
		for i := 0; i < lrm; i++ {
			tmp[i] = m[i][j]
		}
		res = append(res, string(tmp))
	}
	return
}

// ResultOf returns the result of the hex game with the marker of the winner in case there is one.
func ResultOf(s []string) (res string, err error) {
	if MatrixVisit(s, 'X') {
		return "X", nil
	}
	if MatrixVisit(matrixRotate(s), 'O') {
		return "O", nil
	}
	return
}
