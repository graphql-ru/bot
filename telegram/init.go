package telegram

import (
	"log"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

// New prepare initial configuration for telegram api
func New() (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		log.Printf("Unable to authorize telegram client")
		return nil, err
	}

	if os.Getenv("DEBUG") != "" {
		bot.Debug = true
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Client{Bot: bot}, nil
}
