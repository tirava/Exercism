// Package cipher implements a simple shift cipher like Caesar and a more secure substitution cipher.
package cipher

import (
	"strings"
	"unicode"
)

// Caesar is the base struct.
type Caesar struct {
	shift rune
	key   string
}

// NewCaesar returns exemplar of base struct.
func NewCaesar() *Caesar {
	return &Caesar{3, ""}
}

// Encode returns encoded string.
func (c *Caesar) Encode(in string) string {

	outSB := strings.Builder{}
	outSB.Grow(len(in))

	in = strings.ToLower(in)
	//s = unicode.ToLower(s)

	if c.key == "" {
		for _, s := range in {
			if !unicode.IsLetter(s) {
				continue
			}
			// to lower
			s1 := s + c.shift
			if s1 > 'z' {
				s1 -= 26
			} else if s1 < 'a' {
				s1 += 26
			}
			outSB.WriteRune(s1)
		}
	}
	return outSB.String()
}

// Decode returns decoded string.
func (c *Caesar) Decode(in string) string {
	c1 := NewCaesar()
	if c.key == "" {
		c1.shift = -c.shift
	}
	return c1.Encode(in)
}

// NewShift returns shifted distance times string.
func NewShift(distance int) Cipher {
	if distance > 25 || -25 > distance || distance == 0 {
		return nil
	}
	c := NewCaesar()
	c.shift = rune(distance)
	return c
}

// NewVigenere returns shifted key times shifted strings.
func NewVigenere(key string) Cipher {
	c := NewCaesar()
	c.key = key
	return c
}
