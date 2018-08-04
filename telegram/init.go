package telegram

import (
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

// New prepare initial configuration for telegram api
func New() (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		return nil, err
	}

	if os.Getenv("DEBUG") != "" {
		bot.Debug = true
	}

	return &Client{Bot: bot}, nil
}
