// Package dna implements computing - how many times each nucleotide occurs in the string.
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

var dna = "ACGT"

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (h Histogram, err error) {
	h = make(Histogram)
	s := string(d)

	for _, dn := range dna {
		h[dn] = strings.Count(s, string(dn))
		s = strings.ReplaceAll(s, string(dn), "")
	}

	if len(string(s)) > 0 {
		return h, errors.New("incorrect symbol")
	}

	return h, nil
}
