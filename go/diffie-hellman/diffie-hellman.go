// Package diffiehellman implements Diffie-Hellman key exchange.
package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var (
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	_0  = big.NewInt(0)
	_1  = big.NewInt(1)
)

// PrivateKey returns picked private key.
func PrivateKey(p *big.Int) *big.Int {
	ab := _0
	for {
		ab.Rand(rnd, p)
		if ab.Cmp(_1) > 0 {
			break
		}
	}
	return ab
}

// PublicKey returns public keys.
func PublicKey(ab, p *big.Int, g int64) *big.Int {
	AB := _0
	return AB.Exp(big.NewInt(g), ab, p)
}

// SecretKey returns secret key.
func SecretKey(ab, BA, p *big.Int) *big.Int {
	AB := _0
	return AB.Exp(BA, ab, p)
}

// NewPair returns new pairs.
func NewPair(p *big.Int, g int64) (ab, AB *big.Int) {
	ab = PrivateKey(p)
	AB = PublicKey(ab, p, g)
	return
}
