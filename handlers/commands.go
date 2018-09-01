package handlers

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

// Commands handle special bot commands
func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update, next func()) {
	if !update.Message.IsCommand() {
		next()
		return
	}

	log.Printf("[COMMAND] @%s %s", update.Message.From.UserName, update.Message.Text)

	var msg tgbotapi.MessageConfig

	ChatID := update.Message.Chat.ID
	MessageID := update.Message.MessageID

	switch update.Message.Command() {
	case "ping":
		msg = tgbotapi.NewMessage(ChatID, "pong")
		msg.ReplyToMessageID = MessageID
	default:
		msg = tgbotapi.NewMessage(ChatID, "Make @graphql_bot better https://github.com/graphql-ru/bot")
		msg.ReplyToMessageID = MessageID
		msg.DisableWebPagePreview = true
	}

	bot.Send(msg)
	next()
}
