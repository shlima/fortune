package telegram

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCli_HeartBeat(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		setup := MustSetup(t)
		defer setup.ctrl.Finish()

		setup.api.EXPECT().
			Send(gomock.Any()).
			DoAndReturn(func(c Chattable) (Message, error) {
				msg := MustChattableMessage(t, c)
				require.Equal(t, setup.Opts.Channel, msg.ChannelUsername)
				require.Contains(t, msg.Text, "foo")
				return Message{}, nil
			})

		err := setup.HeartBeat("foo")
		require.NoError(t, err)
	})
}

func TestCli_KeyFound(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		setup := MustSetup(t)
		defer setup.ctrl.Finish()

		setup.api.EXPECT().
			Send(gomock.Any()).
			DoAndReturn(func(c Chattable) (Message, error) {
				msg := MustChattableMessage(t, c)
				require.Equal(t, setup.Opts.Channel, msg.ChannelUsername)
				require.Contains(t, msg.Text, "foo")
				return Message{}, nil
			})

		err := setup.KeyFound("foo")
		require.NoError(t, err)
	})
}
