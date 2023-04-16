package pass

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGen_Next(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		t.Parallel()
		gen := New(GenOpts{
			Alphabet: []string{"1", "2"},
			Length:   2,
		})

		expected := []string{
			"11",
			"12",
			"21",
			"22",
		}

		got := gen.All()
		require.Len(t, got, 4)
		require.Len(t, got, int(gen.Permutations()))
		require.Equal(t, int(gen.Permutations()), int(gen.GetIterations()))
		require.ElementsMatch(t, expected, got)

	})

	t.Run("it works", func(t *testing.T) {
		t.Parallel()
		gen := New(GenOpts{
			Alphabet: []string{"0", "1", "2"},
			Length:   3,
		})

		expected := []string{
			"222",
			"221",
			"220",
			"212",
			"211",
			"210",
			"202",
			"201",
			"200",
			"122",
			"121",
			"120",
			"112",
			"111",
			"110",
			"102",
			"101",
			"100",
			"022",
			"021",
			"020",
			"012",
			"011",
			"010",
			"002",
			"001",
			"000",
		}

		got := gen.All()
		require.Len(t, got, 27)
		require.Len(t, got, int(gen.Permutations()))
		require.Equal(t, int(gen.Permutations()), int(gen.iterations)-1)
		require.ElementsMatch(t, expected, got)
	})

	t.Run("it works", func(t *testing.T) {
		t.Parallel()
		gen := New(GenOpts{
			Alphabet: []string{"0", "1", "2"},
			Length:   8,
		})

		got := gen.All()
		require.Len(t, got, int(gen.Permutations()))
		require.Equal(t, int(gen.Permutations()), int(gen.GetIterations()))
	})
}
