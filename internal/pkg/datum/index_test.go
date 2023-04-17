package datum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndex_Random(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		index := Index{"foo": true}
		got, _ := index.Random()
		require.Equal(t, "foo", got)
	})

	t.Run("when empty", func(t *testing.T) {
		index := make(Index)
		got, _ := index.Random()
		require.Equal(t, "", got)
	})
}

func TestIndex_SetTesting(t *testing.T) {
	t.Parallel()

	t.Run("when empty", func(t *testing.T) {
		t.Parallel()

		index := Index{"foo": true}
		got := index.SetTesting("")
		require.Equal(t, index, got)
	})

	t.Run("when not", func(t *testing.T) {
		t.Parallel()

		index := Index{"foo": true}
		got := index.SetTesting("bar")
		require.Equal(t, Index{"bar": true}, got)
	})
}
