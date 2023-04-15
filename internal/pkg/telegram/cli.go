package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Opts struct {
	Token   string
	Channel string
}

type Cli struct {
	Opts
	api IApi
}

func New(opts Opts) (*Cli, error) {
	bot, err := tgbotapi.NewBotAPI(opts.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to init bot: %w", err)
	}

	return &Cli{Opts: opts, api: bot}, nil
}

func (c *Cli) HeartBeat(message string) error {
	msg := tgbotapi.NewMessageToChannel(c.Channel, message)
	_, err := c.api.Send(msg)
	return err
}

func (c *Cli) KeyFound(html string) error {
	msg := tgbotapi.NewMessageToChannel(c.Channel, html)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.DisableWebPagePreview = true
	_, err := c.api.Send(msg)
	return err
}

func (c *Cli) IsReal() bool {
	return true
}
