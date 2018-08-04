package telegram

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// UpdateCallback handle bot updateds
type UpdateCallback func()

// Client wrapper for telegram api
type Client struct {
	Bot *tgbotapi.BotAPI
}

// Start client
func (c *Client) Start() error {
	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	updates, err := c.Bot.GetUpdatesChan(ucfg)

	if err != nil {
		log.Println("Unable to start client")
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("@%s %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Pong")

		c.Bot.Send(msg)
	}

	return nil
}

// onMessage asdasd
func (c *Client) onMessage() {
	log.Println("Hello GraphQl")
}
