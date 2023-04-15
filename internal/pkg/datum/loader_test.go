package datum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		t.Parallel()

		got, err := ReadFile("testdata/uno-dos.txt")
		require.NoError(t, err)
		require.Contains(t, got, "uno")
		require.Contains(t, got, "dos")
	})

	t.Run("it errors", func(t *testing.T) {
		t.Parallel()

		_, err := ReadFile("fake.txt")
		require.Error(t, err)
	})
}

func TestReadFiles(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got, err := ReadFiles("testdata/uno-dos.txt", "testdata/tres.txt")
		require.NoError(t, err)
		require.Contains(t, got, "uno")
		require.Contains(t, got, "dos")
		require.Contains(t, got, "tres")
	})

	t.Run("it errors", func(t *testing.T) {
		t.Parallel()

		_, err := ReadFiles("fake.txt")
		require.Error(t, err)
	})
}
