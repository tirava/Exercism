// Package tournament implements tallying the results of a small football competition.
package tournament

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Tally fills tournament table.
func Tally(r io.Reader, w io.Writer) (err error) {

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r)
	if err != nil {
		return
	}
	s := buf.String()

	//fmt.Println(s)
	lines := strings.Split(s, "\n")
	//fmt.Println(lines)
	_, err = fmt.Fprintf(w, "%-31s|%4s|%4s|%4s|%4s|%4s\n", "Team", "MP ", "W ", "D ", "L ", "P ")

	for _, line := range lines {
		item := strings.Split(line, ";")
		if len(item) > 1 {
			//fmt.Println(item)
			_, err = fmt.Fprintf(w, "%-31s|%4s|%4s|%4s|%4s|%4s\n", item[0], "MP ", "W ", "D ", "L ", "P ")
		}

	}

	return
}
