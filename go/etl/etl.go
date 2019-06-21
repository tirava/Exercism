// Package etl implements etl scoring.
package etl

import "strings"

type Give map[int][]string
type Expect map[string]int

// Transform returns transformated score table.
func Transform(in Give) Expect {
	out := Expect{}
	for i := 1; i <= 10; i++ {
		if _, ok := in[i]; !ok {
			continue
		}
		for _, letter := range in[i] {
			out[strings.ToLower(letter)] = i
		}
	}

	return out
}
