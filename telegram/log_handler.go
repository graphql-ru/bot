package telegram

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// LogHandler log every messages
func LogHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("@%s %s", update.Message.From.UserName, update.Message.Text)
}
