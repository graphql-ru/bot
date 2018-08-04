package telegram

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// Client wrapper for telegram api
type Client struct {
	Bot *tgbotapi.BotAPI
}

// Hello just stays hello
func (c *Client) Hello() {
	log.Println("Hello GraphQl")
}
