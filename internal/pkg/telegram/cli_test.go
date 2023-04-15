package telegram

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shlima/fortune/internal/pkg/bruteforce"
	"github.com/shlima/fortune/internal/pkg/key"
	"github.com/stretchr/testify/require"
)

func Test_addressLink(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := addressLink("foo")
		require.Equal(t, "https://www.blockchain.com/explorer/addresses/btc/foo", got)
	})
}

func TestCli_HeartBeat(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		setup := MustSetup(t)
		defer setup.ctrl.Finish()

		beat := &bruteforce.HeartBit{
			Tried: 1,
			IOps:  2,
		}

		setup.api.EXPECT().
			Send(gomock.Any()).
			DoAndReturn(func(c Chattable) (Message, error) {
				msg := MustChattableMessage(t, c)
				require.Equal(t, setup.Opts.Channel, msg.ChannelUsername)
				require.Contains(t, msg.Text, beat.ToString())
				return Message{}, nil
			})

		err := setup.HeartBeat(beat)
		require.NoError(t, err)
	})
}

func TestCli_KeyFound(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		setup := MustSetup(t)
		defer setup.ctrl.Finish()

		chain := key.Chain{
			Private:      "foo",
			Compressed:   "bar",
			Uncompressed: "baz",
		}

		setup.api.EXPECT().
			Send(gomock.Any()).
			DoAndReturn(func(c Chattable) (Message, error) {
				msg := MustChattableMessage(t, c)
				require.Equal(t, setup.Opts.Channel, msg.ChannelUsername)
				require.Contains(t, msg.Text, chain.Private)
				require.Contains(t, msg.Text, chain.Compressed)
				require.Contains(t, msg.Text, chain.Uncompressed)
				return Message{}, nil
			})

		err := setup.KeyFound(chain)
		require.NoError(t, err)
	})
}
