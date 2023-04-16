package blockchain

import (
	"context"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestCli_Addresses(t *testing.T) {
	t.Cleanup(httpmock.DeactivateAndReset)
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		cli := MustSetup(t)

		cli.ExpectJSONGet(
			"https://blockchain.info/multiaddr?active=1AxSQFHqQ2nFUbZwcWSmHYrYumuQnK2nYG%7C1PyRAXKXT9dPXq6A1rLLAMajQgR37KjEDa&n=1",
			200,
			MustTextFixture(t, "testdata/01.json"))

		got, err := cli.Addresses(context.Background(), []string{"1AxSQFHqQ2nFUbZwcWSmHYrYumuQnK2nYG", "1PyRAXKXT9dPXq6A1rLLAMajQgR37KjEDa"})
		require.NoError(t, err)
		require.Len(t, got, 2)

		require.Equal(t, Address{
			Address:       "1AxSQFHqQ2nFUbZwcWSmHYrYumuQnK2nYG",
			NTx:           14,
			TotalReceived: 101058761,
			TotalSent:     101058761,
		}, got[0])

		require.Equal(t, Address{
			Address: "1PyRAXKXT9dPXq6A1rLLAMajQgR37KjEDa",
		}, got[1])
	})
}
