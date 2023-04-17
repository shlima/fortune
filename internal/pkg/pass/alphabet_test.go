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
