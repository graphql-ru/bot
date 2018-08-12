package handlers

import (
	"gopkg.in/telegram-bot-api.v4"
)

// Commands handle special bot commands
func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update, next func()) {
	if !update.Message.IsCommand() {
		next()
		return
	}

	var msg tgbotapi.MessageConfig

	switch update.Message.Command() {
	case "ping":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "pong")
		msg.ReplyToMessageID = update.Message.MessageID
	default:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Make @graphql_bot better https://github.com/graphql-ru/bot")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.DisableWebPagePreview = true
	}

	bot.Send(msg)
	next()
}
