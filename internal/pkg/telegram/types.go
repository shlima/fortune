package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//go:generate mockgen -source types.go -destination ../../mock/telegram.go -package mock -mock_names ICli=MockTelegramCli,IApi=MockTelegramApi

type (
	MessageConfig = tgbotapi.MessageConfig
	Chattable     = tgbotapi.Chattable
	Message       = tgbotapi.Message
)

type ICli interface {
	HeartBeat(message string) error
	KeyFound(message string) error
	IsReal() bool
}

// IApi should be implemented by tgbotapi.BotAPI
type IApi interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}
