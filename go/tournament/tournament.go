// Package tournament implements tallying the results of a small football competition.
package tournament

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	win = iota
	draw
	loss
	command1 = 0
	command2 = 1
)

// Tally fills tournament table.
func Tally(r io.Reader, w io.Writer) (err error) {

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r)
	//if err != nil {
	//	return
	//}
	s := buf.String()

	//results := map[string][3]int{} // [3]int{won, drawn, lost}
	results := make(map[string][3]int) // [3]int{won, drawn, lost}
	formatHeader := "%-31s|%4s|%4s|%4s|%4s|%4s\n"
	formatBody := "%-31s|%4s|%3d |%3d |%3d |%4s\n"

	lines := strings.Split(s, "\n")
	_, _ = fmt.Fprintf(w, formatHeader, "Team", "MP ", "W ", "D ", "L ", "P ")

	for _, line := range lines {
		if len(line) <= 1 {
			continue
		}
		result := strings.Split(line, ";")

		switch result[2] { // win, draw, loss
		case "win":
			c1 := results[result[command1]]
			c1[win]++
			results[result[command1]] = c1
		//	results[result[command1]][win]++
		//case "draw":
		//	results[result[command1]][draw]++
		//	results[result[command2]][draw]++
		//case "loss":
		//	results[result[command2]][loss]++
		default:
			//x := results["www"]
			//x[0]++
			//results["www"] = x
			//fmt.Println(results)
		}

		for k, v := range results {
			_, _ = fmt.Fprintf(w, formatBody, k, "MP ", v[win], v[draw], v[loss], "P ")
		}

	}

	return
}
