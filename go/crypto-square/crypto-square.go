// Package cryptosquare implements encoding text.
package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

var r = regexp.MustCompile("[^0-9a-zA-Z]+")

// Encode returns encoded version of plain text.
func Encode(pt string) (ct string) {
	ct = r.ReplaceAllString(pt, "")
	ct = strings.ToLower(ct)
	q := int(math.Sqrt(float64(len(ct))))
	c := q + 1
	r := len(ct) / q
	ll := c - c*raaaaaaaaaaaaaaaaaa
	sl := make([]string, 0)
	for i := 0; i < q; i++ {
		//sl = append(sl, ct[i*(q+1) : 2*i*(q+1)])
	}
	println(c, r, ll, sl)
	return
}
