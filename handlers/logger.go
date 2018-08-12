package handlers

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// Logger log every messages
func Logger(bot *tgbotapi.BotAPI, update tgbotapi.Update, next func()) {
	if update.Message.Text == "" {
		next()
		return
	}

	log.Printf("@%s %s", update.Message.From.UserName, update.Message.Text)
	next()
}
