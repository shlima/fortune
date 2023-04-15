package key

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/require"
)

func TestGenerator_SetTesting(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		gen := New()
		require.Empty(t, gen.testing)

		address := "foo"
		gen.SetTesting(address)
		require.Equal(t, gen.testing, address)
	})
}

func TestGenerator_Generate(t *testing.T) {
	t.Parallel()

	t.Run("they differ", func(t *testing.T) {
		t.Parallel()

		got01, err01 := New().Generate()
		got02, err02 := New().Generate()
		require.NoError(t, err01)
		require.NoError(t, err02)

		require.NotEqual(t, got01, got02)
	})

	t.Run("it works", func(t *testing.T) {
		t.Parallel()

		got, err := New().Generate()
		require.NoError(t, err)
		require.Len(t, MustHexDecode(t, got.Private), 256>>3)

		_, err = btcutil.DecodeAddress(got.Compressed, &chaincfg.MainNetParams)
		require.NoError(t, err)
		_, err = btcutil.DecodeAddress(got.Uncompressed, &chaincfg.MainNetParams)
		require.NoError(t, err)
	})
}

func MustHexDecode(t *testing.T, input string) []byte {
	got, err := hex.DecodeString(input)
	require.NoError(t, err)
	return got
}
