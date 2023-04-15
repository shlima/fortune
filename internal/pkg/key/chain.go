package key

import (
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
		Private:      "Private",
		Compressed:   "Compressed",
		Uncompressed: "Uncompressed",
	}

	switch {
	case time.Now().Nanosecond() > 499999999:
		out.Compressed = address
	default:
		out.Uncompressed = address
	}

	return out
}
