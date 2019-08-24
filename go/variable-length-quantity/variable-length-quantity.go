// Package variablelengthquantity implements variable length quantity encoding and decoding.
package variablelengthquantity

import "errors"

// DecodeVarint returns decoded VLQ.
func DecodeVarint(input []byte) (out []uint32, err error) {

	var outI uint64

	for i, in := range input {
		flush := false
		if i == len(input)-1 && in&0x80 == 0x80 {
			err = errors.New("last byte incomplete")
			return
		}
		if in&0x80 == 0 {
			flush = true
		}
		in &= 0x7f
		in <<= 1
		outI <<= 8
		outI |= uint64(in)
		outI >>= 1

		if flush {
			out = append(out, uint32(outI))
			outI = 0
		}
	}

	return
}

// EncodeVarint returns encoded VLQ.
func EncodeVarint(input []uint32) (out []byte) {

	for _, in := range input {
		outB := make([]byte, 0, 5)
		for i := uint(0); i < 5; i++ {
			ror := uint64((uint64(in) << i) >> (i * 8))
			rolB := byte(ror)
			if rolB == 0 && i > 0 && ror == 0 {
				break
			}
			if i == 0 {
				rolB &= 0x7f
			} else {
				rolB |= 0x80
			}
			outB = append([]byte{rolB}, outB...)
		}
		out = append(out, outB...)
	}

	return out
}
