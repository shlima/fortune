package key

import (
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

func TestGenerator_BrainSHA256(t *testing.T) {
	t.Parallel()

	// @refs https://www.bitaddress.org/bitaddress.org-v3.3.0-SHA256-dec17c07685e1870960903d8f58090475b25af946fe95a734f88408cef4aa194.html
	t.Run("it works", func(t *testing.T) {
		got, err := New().BrainSHA256([]byte("example of brain wallet"))
		require.NoError(t, err)
		require.Equal(t, "17QZasmw4MKwoUhmTiqB73V7Mom3WNqVNR", got.Uncompressed)

		wif, err := got.ToWIF()
		require.NoError(t, err)
		require.Equal(t, "5KWv77NznJsh8NCGpjYXtVyqd26h1sZ95VTeWkwG57TQS6DRsgz", wif)
	})
}
