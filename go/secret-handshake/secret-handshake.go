// Package secret implements binary handshake.
package secret

var words = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

// Handshake returns slice of handshake words
func Handshake(code uint) (h []string) {

	for i := uint(1); i <= 8; i <<= 1 {
		if code&(i|16) > 16 {
			h = append([]string{words[i]}, h...)
		} else if code&i > 0 {
			h = append(h, words[i])
		}
	}

	return
}
