package bruteforce

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeartBit_ToString(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		bit := HeartBit{
			Tried: 1,
			IOps:  2,
		}

		got := bit.ToString()
		require.Equal(t, "tried 1 keys (2 ops/sec)", got)
	})
}
