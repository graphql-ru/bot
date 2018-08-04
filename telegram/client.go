package telegram

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// UpdateHandler handle bot updateds
type UpdateHandler func(bot *tgbotapi.BotAPI, update tgbotapi.Update, next func())

// Client wrapper for telegram api
type Client struct {
	Bot      *tgbotapi.BotAPI
	handlers []UpdateHandler
}

// Start client
func (c *Client) Start() error {
	ucfg := tgbotapi.NewUpdate(1)
	ucfg.Timeout = 60

	goNext := false
	next := func() {
		goNext = true
	}

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
			goNext = false

			handler(c.Bot, update, next)

			if !goNext {
				break
			}
		}
	}

	return nil
}

// Use adds handler to handlers array
func (c *Client) Use(handler UpdateHandler) {
	c.handlers = append(c.handlers, handler)
}
