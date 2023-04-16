package blockchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddress_FinalBalanceAmount(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		a := Address{FinalBalance: 100}
		got := a.FinalBalanceAmount()
		require.Equal(t, "0.000001 BTC", got)
	})
}
