package handlers

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// Logger log every messages
func Logger(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("@%s %s", update.Message.From.UserName, update.Message.Text)
}
