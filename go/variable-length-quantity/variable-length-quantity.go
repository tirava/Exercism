// Package variablelengthquantity implements variable length quantity encoding and decoding.
package variablelengthquantity

// DecodeVarint returns decoded VLQ.
func DecodeVarint(input []byte) (out []uint32, err error) {

	//for _, b := range input {
	//
	//}

	return nil, nil
}

// EncodeVarint returns encoded VLQ.
func EncodeVarint(input []uint32) (out []byte) {

	for _, in := range input {
		outB := make([]byte, 0, 5)
		for i := uint32(0); i < 5; i++ {
			rol := uint64((uint64(in) << i) >> (i * 8))
			rolB := byte(rol)
			if rolB == 0 && i > 0 && rol == 0 {
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
