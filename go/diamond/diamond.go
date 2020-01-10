// Package diamond implements the diamond kata.
package diamond

import (
	"errors"
	"fmt"
	"strings"
)

// Gen returns diamond.
func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", errors.New("invalid letter")
	}

	dia := strings.Builder{}

	for i := byte('A'); i < b; i++ {
		dia.WriteString(writeLine(i, b))
	}
	for i := b; i >= 'A'; i-- {
		dia.WriteString(writeLine(i, b))
	}

	return dia.String(), nil
}

func writeLine(b, letters byte) string {
	s1 := strings.Repeat(" ", int(letters-b))
	s2 := strings.Repeat(" ", int((b-'A')*2-1))
	bs := ""

	if b == 'A' {
		bs = string(b)
	} else {
		bs = fmt.Sprintf("%s%s%s", string(b), s2, string(b))
	}

	return fmt.Sprintf("%s%s%s\n", s1, bs, s1)
}
