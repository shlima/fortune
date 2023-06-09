package key

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type KeyChain struct {
	Private      string
	Compressed   string
	Uncompressed string
}

func NewTestingKeyChain(address string) KeyChain {
	out := KeyChain{
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

func (k *KeyChain) ToString() string {
	return fmt.Sprintf(
		"Private: %s Compressed: %s Ucomprssed: %s",
		k.Private, k.Compressed, k.Uncompressed)
}

func (k *KeyChain) ToWIF() (string, error) {
	bytea, err := hex.DecodeString(k.Private)
	if err != nil {
		return "", fmt.Errorf("failed to decode: %w", err)
	}

	return PrivToWIF(bytea)
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
