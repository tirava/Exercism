// Package variablelengthquantity implements variable length quantity encoding and decoding.
package variablelengthquantity

import "errors"

// DecodeVarint returns decoded VLQ.
func DecodeVarint(input []byte) (out []uint32, err error) {
	var outI uint64

	for i, in := range input {
		if i == len(input)-1 && in&0x80 == 0x80 {
			err = errors.New("last byte incomplete")
			return
		}
		outI = ((outI << 8) | uint64((in&0x7f)<<1)) >> 1
		if in&0x80 == 0 {
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
			rorB := byte(ror)
			if rorB == 0 && i > 0 && ror == 0 {
				break
			}
			rorB |= 0x80
			if i == 0 {
				rorB &= 0x7f
			}
			outB = append([]byte{rorB}, outB...)
		}
		out = append(out, outB...)
	}

	return out
}
