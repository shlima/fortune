package key

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func MustHexDecode(t *testing.T, input string) []byte {
	got, err := hex.DecodeString(input)
	require.NoError(t, err)
	return got
}
