package telegram

import (
	"fmt"
	"github.com/shlima/fortune/internal/pkg/key"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shlima/fortune/internal/pkg/bruteforce"
)

type Opts struct {
	Token   string
	Channel string
}

type Cli struct {
	Opts
	bot *tgbotapi.BotAPI
}

func New(opts Opts) (*Cli, error) {
	bot, err := tgbotapi.NewBotAPI(opts.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to init bot: %w", err)
	}

	return &Cli{Opts: opts, bot: bot}, nil
}

func (c *Cli) SendHeartBeat(heartbit *bruteforce.HeartBit) error {
	msg := tgbotapi.NewMessageToChannel(c.Channel, heartbit.ToString())
	_, err := c.bot.Send(msg)
	return err
}

func (c *Cli) KeyFound(chain key.Chain) error {
	msg := tgbotapi.NewMessageToChannel(c.Channel, chain.ToString())
	_, err := c.bot.Send(msg)
	return err
}
