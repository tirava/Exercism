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
func PublicKey(ab, p *big.Int, g int64) *big.Int {
	return ab.Exp(big.NewInt(g), ab, p)
}

// SecretKey returns secret key.
func SecretKey(ab, BA, p *big.Int) *big.Int {
	//s = B**a mod p
	//s = A**b mod p
	AB := big.NewInt(0)
	AB.Exp(BA, ab, p)
	return AB
}

// NewPair returns new pairs.
func NewPair(p *big.Int, g int64) (ab, AB *big.Int) {
	ab = PrivateKey(p)
	AB = PublicKey(ab, p, g)
	return
}
