package key

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type Generator struct {
	max     *big.Int
	testing string // test address
}

func New() *Generator {
	return &Generator{
		// Max value, a 256-bits integer, i.e 2^256 - 1
		max: big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil),
	}
}

func (g *Generator) SetTesting(address string) IGenerator {
	g.testing = address
	return g
}

// Generate
// private key: A secret number, known only to the person that generated it.
// A private key is essentially a randomly generated number.
//
// In Bitcoin, someone with the private key that corresponds to funds on the block
// chain can spend the funds.
//
// In Bitcoin, a private key is a single unsigned 256 bit integer (32 bytes).
func (g *Generator) Generate() (out KeyChain, err error) {
	if g.testing != "" {
		return NewTestingKeyChain(g.testing), nil
	}

	res, err := rand.Int(rand.Reader, g.max)
	if err != nil {
		return out, fmt.Errorf("failed to rand: %w", err)
	}

	return KeyChainFromPriv(res.Bytes())
}

// BrainSHA256 generates a brain address base on SHA256(passphrase)
func (g *Generator) BrainSHA256(password []byte) (out KeyChain, err error) {
	if g.testing != "" {
		return NewTestingKeyChain(g.testing), nil
	}

	hash := sha256.Sum256(password)
	return KeyChainFromPriv(hash[:])
}
