package telegram

import (
	"os"

	tg "gopkg.in/telegram-bot-api.v4"
)

// New prepare initial configuration for telegram api
func New() (*Client, error) {
	bot, err := tg.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		return nil, err
	}

	return &Client{Bot: bot}, nil
}
