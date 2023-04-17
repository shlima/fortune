package pass

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshallState(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := MarshallState([]int{})
		require.Equal(t, "", got)

		got = MarshallState([]int{1, 2})
		require.Equal(t, "1,2", got)
	})
}

func TestUnmarshalState(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got, err := UnmarshalState("")
		require.NoError(t, err)
		require.Equal(t, []int{}, got)

		got, err = UnmarshalState("1,2")
		require.NoError(t, err)
		require.Equal(t, []int{1, 2}, got)
	})
}

func TestMarshallAlphabet(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := MarshallAlphabet([]string{"a", "b"})
		require.Equal(t, "ab", got)
	})
}
