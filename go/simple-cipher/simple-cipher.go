// Package cipher implements a simple shift cipher like Caesar and a more secure substitution cipher.
package cipher

// Caesar is the base struct.
type Caesar struct {
}

// NewCaesar returns exemplar of base struct.
func NewCaesar() *Caesar {
	return &Caesar{}
}

// Encode returns encoded string.
func (c *Caesar) Encode(in string) string {
	return ""
}

// Decode returns decoded string.
func (c *Caesar) Decode(in string) string {
	return ""
}

// NewShift returns shifted distance times string.
func NewShift(distance int) *Caesar {
	return &Caesar{}
}

// NewVigenere returns shifted key times shifted strings.
func NewVigenere(key string) *Caesar {
	return &Caesar{}
}
