package key

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChain_ToString(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		chain := Chain{
			Private:      "Private01",
			Compressed:   "Compressed02",
			Uncompressed: "Uncompressed02",
		}

		got := chain.ToString()
		require.Contains(t, got, chain.Private)
		require.Contains(t, got, chain.Compressed)
		require.Contains(t, got, chain.Uncompressed)
	})
}

func TestNewTestingChain(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		address := "foo"
		got := NewTestingChain(address)
		require.NotEmpty(t, got.Private)
		require.NotEmpty(t, got.Compressed)
		require.NotEmpty(t, got.Uncompressed)
		require.Contains(t, []string{got.Compressed, got.Uncompressed}, address)
	})
}
