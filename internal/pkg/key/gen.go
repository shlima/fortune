package key

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
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
func (g *Generator) Generate() (out Chain, err error) {
	if g.testing != "" {
		return NewTestingChain(g.testing), nil
	}

	res, err := rand.Int(rand.Reader, g.max)
	if err != nil {
		return out, fmt.Errorf("failed to rand: %w", err)
	}

	bytea := res.Bytes()
	_, pub := btcec.PrivKeyFromBytes(bytea)
	compressed, err := btcutil.NewAddressPubKey(pub.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return out, fmt.Errorf("failed to compress key: %w", err)
	}

	uncompressed, err := btcutil.NewAddressPubKey(pub.SerializeUncompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return out, fmt.Errorf("failed to uncompress key")
	}

	return Chain{
		Private:      hex.EncodeToString(bytea),
		Compressed:   compressed.EncodeAddress(),
		Uncompressed: uncompressed.EncodeAddress(),
	}, nil
}
