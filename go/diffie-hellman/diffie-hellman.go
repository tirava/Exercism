// Package diffiehellman implements Diffie-Hellman key exchange.
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

var _1 = big.NewInt(1)
var _23 = big.NewInt(23)

// PrivateKey returns picked private key.
func PrivateKey(p *big.Int) (ab *big.Int) {
	for {
		ab, _ = rand.Int(rand.Reader, p)
		if ab.Cmp(_1) > 0 {
			break
		}
	}
	return
}

// PublicKey returns public keys.
func PublicKey(ab, p *big.Int, g int64) (AB *big.Int) {

	return
}

// SecretKey returns secret key.
func SecretKey(ab, AB, p *big.Int) (s *big.Int) {

	return
}

// NewPair returns new pairs.
func NewPair(p *big.Int, g int64) (ab, AB *big.Int) {

	return
}
