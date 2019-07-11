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
	lenCT := len(ct)
	//q := int(math.Sqrt(float64(lenCT)))
	//cols, rows := q, q
	//if lenCT > q*q {
	//	//fmt.Println(lenCT, q*2)
	//	cols++
	//}
	//if cols*rows < lenCT {
	//	//fmt.Println(lenCT, q*2)
	//	cols++
	//}
	//ll := cols - (cols*rows - lenCT)
	//fmt.Println(cols, rows, ll, lenCT, q, ct)
	//sl := make([]string, 0)
	//spaces := ""
	//lenCol := cols
	//for i := 0; i < rows; i++ {
	//	if i == rows-1 {
	//		lenCol = ll
	//		//fmt.Println(cols, rows, ll, lenCT, "repeat", cols-ll)
	//		spaces = strings.Repeat(" ", cols-ll)
	//	}
	//	sl = append(sl, ct[i*cols:i*cols+lenCol]+spaces)
	//	//fmt.Println(sl)
	//}
	//ct = ""
	//for i := 0; i < cols; i++ {
	//	for _, ss := range sl {
	//		ct += string(ss[i])
	//	}
	//	if i < cols-1 {
	//		ct += " "
	//	}
	//}
	////fmt.Println(ct)
	//fmt.Println(cols, rows, ll, sl, q, "->", ct)
	cols := int(math.Ceil(math.Sqrt(float64(lenCT))))
	rows := int(math.Ceil(float64(lenCT) / float64(cols)))
	enc := ""
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if j*cols+i < lenCT {
				enc += string(ct[j*cols+i])
			} else {
				enc += " "
			}
		}
		if len(enc) < cols*rows {
			enc += " "
		}
	}

	return enc
}
