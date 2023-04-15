package key

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type Chain struct {
	Private      string
	Compressed   string
	Uncompressed string
}

func (c *Chain) ToString() string {
	return fmt.Sprintf(
		"Private: %s Compressed: %s Ucomprssed: %s",
		c.Private, c.Compressed, c.Uncompressed)
}

func NewTestingChain(address string) Chain {
	out := Chain{
		Private:      randomHex(32),
		Compressed:   randomHex(14),
		Uncompressed: randomHex(14),
	}

	switch {
	case time.Now().Nanosecond() > 499999999:
		out.Compressed = address
	default:
		out.Uncompressed = address
	}

	return out
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
