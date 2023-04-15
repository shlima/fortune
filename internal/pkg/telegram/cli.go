package telegram

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shlima/fortune/internal/pkg/bruteforce"
	"github.com/shlima/fortune/internal/pkg/key"
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

func (c *Cli) HeartBeat(heartbit *bruteforce.HeartBit) error {
	msg := tgbotapi.NewMessageToChannel(c.Channel, heartbit.ToString())
	_, err := c.api.Send(msg)
	return err
}

func (c *Cli) KeyFound(chain key.Chain) error {
	msg := tgbotapi.NewMessageToChannel(c.Channel, chainMessage(chain))
	msg.ParseMode = tgbotapi.ModeHTML
	msg.DisableWebPagePreview = true
	_, err := c.api.Send(msg)
	return err
}

func chainMessage(chain key.Chain) string {
	b := new(strings.Builder)
	b.WriteString("💰")
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("<b>Private</b>: %s", chain.Private))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf(`<b>Compressed</b>: <a href="%s">%s</a>`, addressLink(chain.Compressed), chain.Compressed))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf(`<b>Uncompressed</b>: <a href="%s">%s</a>`, addressLink(chain.Uncompressed), chain.Uncompressed))

	return b.String()
}

func addressLink(address string) string {
	return fmt.Sprintf("https://www.blockchain.com/explorer/addresses/btc/%s", address)
}
