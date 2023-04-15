package telegram

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNoOp_IsReal(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := NewNoOp().IsReal()
		require.False(t, got)
	})
}
