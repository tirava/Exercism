// Package cryptosquare implements encoding text.
package cryptosquare

import (
	"regexp"
	"strings"
)

var r = regexp.MustCompile("[^0-9a-zA-Z]+")

// Encode returns encoded version of plain text.
func Encode(pt string) (ct string) {
	ct = r.ReplaceAllString(pt, "")
	ct = strings.ToLower(ct)
	//println(ct)
	return
}
