// Package diffiehellman implements Diffie-Hellman key exchange.
package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var (
	rnd    = rand.New(rand.NewSource(time.Now().UnixNano()))
	_0, _1 = big.NewInt(0), big.NewInt(1)
)

// PrivateKey returns picked private key.
func PrivateKey(p *big.Int) *big.Int {
	for {
		_0.Rand(rnd, p)
		if _0.Cmp(_1) > 0 {
			break
		}
	}
	return _0
}

// PublicKey returns public keys.
func PublicKey(ab, p *big.Int, g int64) *big.Int {
	return _0.Exp(big.NewInt(g), ab, p)
}

// SecretKey returns secret key.
func SecretKey(ab, BA, p *big.Int) *big.Int {
	return _0.Exp(BA, ab, p)
}

// NewPair returns new pairs.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	ab := PrivateKey(p)
	return ab, PublicKey(ab, p, g)
}
