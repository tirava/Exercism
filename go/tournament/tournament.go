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
	command1 = iota
	command2
	played
)

type score struct {
	win, draw, loss int
}

type result map[string]*score

type resultPair struct {
	command string
	points  int
}

type byPoints []resultPair

// Tally fills tournament table.
func Tally(r io.Reader, w io.Writer) (err error) {

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r)
	//if err != nil {
	//	return
	//}
	s := buf.String()

	results := make(result)
	formatHeader := "%-31s|%4s|%4s|%4s|%4s|%3s\n"
	formatBody := "%-31s|%3d |%3d |%3d |%3d |%3d\n"

	lines := strings.Split(s, "\n")
	_, _ = fmt.Fprintf(w, formatHeader, "Team", "MP ", "W ", "D ", "L ", "P")

	for _, line := range lines {
		if len(line) <= 1 || line[0] == '#' {
			continue
		}
		result := strings.Split(line, ";")
		if len(result) != 3 {
			return errors.New("incorrect commands")
		}

		if _, ok := results[result[command1]]; !ok {
			results[result[command1]] = &score{}
		}
		if _, ok := results[result[command2]]; !ok {
			results[result[command2]] = &score{}
		}

		switch result[played] {
		case "win":
			results[result[command1]].win++
			results[result[command2]].loss++
		case "draw":
			results[result[command1]].draw++
			results[result[command2]].draw++
		case "loss":
			results[result[command1]].loss++
			results[result[command2]].win++
		default:
			return errors.New("incorrect result")
		}
	}

	pl := make(byPoints, len(results))
	i := 0
	for k, v := range results {
		pl[i] = resultPair{k, v.win*3 + v.draw}
		i++
	}
	//sort.Sort(pl)
	sort.Sort(sort.Reverse(pl))

	for _, v := range pl {
		_, _ = fmt.Fprintf(w, formatBody, v.command,
			results[v.command].win+results[v.command].draw+results[v.command].loss,
			results[v.command].win, results[v.command].draw, results[v.command].loss, v.points)
	}

	return
}

func (p byPoints) Len() int { return len(p) }

func (p byPoints) Less(i, j int) bool {
	if p[i].points == p[j].points {
		return p[i].command > p[j].command
	} else {
		return p[i].points < p[j].points
	}
}

func (p byPoints) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
