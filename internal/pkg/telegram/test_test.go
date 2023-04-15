package telegram

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shlima/fortune/internal/mock"
)

type Setup struct {
	ctrl *gomock.Controller
	api  *mock.MockTelegramApi
	*Cli
}

func MustSetup(t *testing.T) *Setup {
	ctrl := gomock.NewController(t)
	api := mock.NewMockTelegramApi(ctrl)

	return &Setup{
		ctrl: ctrl,
		api:  api,
		Cli: &Cli{
			api: api,
			Opts: Opts{
				Channel: "Channel",
				Token:   "Token",
			},
		},
	}
}

func MustChattableMessage(t *testing.T, in Chattable) MessageConfig {
	msg, ok := in.(MessageConfig)
	require.True(t, ok)
	return msg
}
