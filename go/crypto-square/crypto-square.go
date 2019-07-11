// Package cryptosquare implements encoding text.
package cryptosquare

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

var r = regexp.MustCompile("[^0-9a-zA-Z]+")

// Encode returns encoded version of plain text.
func Encode(pt string) (ct string) {
	ct = r.ReplaceAllString(pt, "")
	ct = strings.ToLower(ct)
	lenCT := len(ct)
	q := int(math.Sqrt(float64(lenCT)))
	cols, rows := q, q
	if lenCT > q*2 {
		cols++
	}
	ll := cols - (cols*rows - lenCT)
	sl := make([]string, 0)
	spaces := ""
	lenCol := cols
	for i := 0; i < rows; i++ {
		if i == rows-1 {
			lenCol = ll
			spaces = strings.Repeat(" ", cols-ll)
		}
		sl = append(sl, ct[i*cols:i*cols+lenCol]+spaces)
		//fmt.Println(sl)
	}
	s := ""
	for i := 0; i < cols; i++ {
		for _, ss := range sl {
			s += string(ss[i])
		}
	}
	fmt.Println(s)
	fmt.Println(cols, rows, ll, sl)
	return
}
