package telegram

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// UpdateHandler handle bot updateds
type UpdateHandler func(bot *tgbotapi.BotAPI, update tgbotapi.Update)

// Client wrapper for telegram api
type Client struct {
	Bot      *tgbotapi.BotAPI
	handlers []UpdateHandler
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

		for _, handler := range c.handlers {
			handler(c.Bot, update)
		}
	}

	return nil
}

// Use calls registerd handlers on any update
func (c *Client) Use(handler UpdateHandler) {
	c.handlers = append(c.handlers, handler)
}
