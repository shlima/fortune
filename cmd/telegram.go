package cmd

import (
	"fmt"
	"time"

	telegram2 "github.com/shlima/fortune/internal/pkg/telegram"
	"github.com/urfave/cli/v2"
)

var FlagTelegramToken = &cli.StringFlag{
	Name:    "telegram-token",
	Usage:   "token of the telegram bot",
	EnvVars: []string{"TELEGRAM_TOKEN"},
}

var FlagTelegramChannel = &cli.StringFlag{
	Name:    "telegram-channel",
	Usage:   "@channel name for the notifications (bot should be added as an administrator)",
	EnvVars: []string{"TELEGRAM_CHANNEL"},
}

var FlagTelegramPingSec = &cli.IntFlag{
	Name:    "telegram-ping-sec",
	Usage:   "send status each N seconds to telegram",
	EnvVars: []string{"TELEGRAM_PING_SEC"},
	Value:   int(time.Hour.Seconds()),
}

func NewTelegram(c *cli.Context) telegram2.ICli {
	token := c.String(FlagTelegramToken.Name)
	channel := c.String(FlagTelegramChannel.Name)
	if token == "" {
		return telegram2.NewNoOp()
	}

	client, err := telegram2.New(telegram2.Opts{Token: token, Channel: channel})
	if err != nil {
		panic(fmt.Errorf("failed to init telegram: %w", err))
	}

	return client
}
