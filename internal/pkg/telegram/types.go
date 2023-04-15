package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shlima/fortune/internal/pkg/bruteforce"
	"github.com/shlima/fortune/internal/pkg/key"
)

//go:generate mockgen -source types.go -destination ../../mock/telegram.go -package mock -mock_names ICli=MockTelegramCli,IApi=MockTelegramApi

type (
	MessageConfig = tgbotapi.MessageConfig
	Chattable     = tgbotapi.Chattable
	Message       = tgbotapi.Message
)

type ICli interface {
	HeartBeat(heartbit *bruteforce.HeartBit) error
	KeyFound(chain key.Chain) error
}

// IApi should be implemented by tgbotapi.BotAPI
type IApi interface {
	Send(c tgbotapi.Chattable) (Message, error)
}
