package pass

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseAlphabets(t *testing.T) {
	t.Parallel()

	t.Run("it works with dictionary", func(t *testing.T) {
		t.Parallel()
		got := ParseAlphabets([]string{"digits"})
		require.Equal(t, Digits, got)
	})

	t.Run("it works with dictionary and custom", func(t *testing.T) {
		t.Parallel()
		got := ParseAlphabets([]string{"digits", "a"})
		require.Contains(t, got, "1")
		require.Contains(t, got, "a")
	})

	t.Run("it uniques", func(t *testing.T) {
		t.Parallel()
		got := ParseAlphabets([]string{"aab"})
		require.Equal(t, []string{"a", "b"}, got)
	})

	t.Run("it works with custom", func(t *testing.T) {
		t.Parallel()
		got := ParseAlphabets([]string{"a,1"})
		require.Equal(t, []string{"a", ",", "1"}, got)
	})
}

func TestShuffleIfNonZero(t *testing.T) {
	t.Parallel()

	t.Run("it do nothing if zero", func(t *testing.T) {
		t.Parallel()

		in := EnglishUpper
		got := ShuffleIfNonZero(in, 0)
		require.Equal(t, got, in)
	})

	t.Run("it works with the same seed", func(t *testing.T) {
		t.Parallel()

		in := EnglishUpper
		got01 := ShuffleIfNonZero(in, 1)
		got02 := ShuffleIfNonZero(in, 1)

		require.NotEqual(t, got01, in)
		require.NotEqual(t, got02, in)

		require.Equal(t, got01, got02)
	})
}
