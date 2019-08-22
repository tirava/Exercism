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

	var s1 rune
	var j int
	for _, s := range in {
		if !unicode.IsLetter(s) {
			continue
		}
		s = unicode.ToLower(s)
		if c.key == "" {
			s1 = s + c.shift
		} else {
			s1 = s + (rune(c.key[j%len(c.key)])-'a')*c.shift
		}
		if s1 > 'z' {
			s1 -= 26
		} else if s1 < 'a' {
			s1 += 26
		}
		outSB.WriteRune(s1)
		j++
	}

	return outSB.String()
}

// Decode returns decoded string.
func (c *Caesar) Decode(in string) string {
	c1 := NewCaesar()
	if c.key == "" {
		c1.shift = -c.shift
	} else {
		c1.key = c.key
		c1.shift = -1
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
	if len(key) == 0 {
		return nil
	}
	a := 0
	for _, s := range key {
		if !unicode.IsLetter(s) || unicode.IsUpper(s) {
			return nil
		}
		if s == 'a' {
			a++
		}
	}
	if a == len(key) {
		return nil
	}
	c := NewCaesar()
	c.key = key
	c.shift = 1
	return c
}
