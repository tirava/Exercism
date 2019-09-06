// Package tournament implements tallying the results of a small football competition.
package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	player1 = iota
	player2
	played
)

type score struct {
	win, draw, loss int
}

type result map[string]*score

type resultPair struct {
	player string
	points int
}

type byPoints []resultPair

// Tally fills tournament table.
func Tally(r io.Reader, w io.Writer) (err error) {

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r)
	if err != nil {
		return
	}
	s := buf.String()

	results := make(result)
	formatHeader := "%-31s|%4s|%4s|%4s|%4s|%3s\n"
	formatBody := "%-31s|%3d |%3d |%3d |%3d |%3d\n"

	lines := strings.Split(s, "\n")
	fmt.Fprintf(w, formatHeader, "Team", "MP ", "W ", "D ", "L ", "P")

	for _, line := range lines {
		if len(line) <= 1 || line[0] == '#' {
			continue
		}
		result := strings.Split(line, ";")
		if len(result) != 3 {
			return errors.New("incorrect commands")
		}

		if _, ok := results[result[player1]]; !ok {
			results[result[player1]] = &score{}
		}
		if _, ok := results[result[player2]]; !ok {
			results[result[player2]] = &score{}
		}

		switch result[played] {
		case "win":
			results[result[player1]].win++
			results[result[player2]].loss++
		case "draw":
			results[result[player1]].draw++
			results[result[player2]].draw++
		case "loss":
			results[result[player1]].loss++
			results[result[player2]].win++
		default:
			return errors.New("incorrect result")
		}
	}

	bp := make(byPoints, len(results))
	i := 0
	for k, v := range results {
		bp[i] = resultPair{k, v.win*3 + v.draw}
		i++
	}
	sort.Slice(bp, func(i, j int) bool {
		if bp[i].points == bp[j].points {
			return bp[i].player < bp[j].player
		}
		return bp[i].points > bp[j].points
	})

	for _, v := range bp {
		fmt.Fprintf(w, formatBody, v.player,
			results[v.player].win+results[v.player].draw+results[v.player].loss,
			results[v.player].win, results[v.player].draw, results[v.player].loss, v.points)
	}

	return
}
