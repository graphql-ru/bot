package handlers

import (
	"gopkg.in/telegram-bot-api.v4"
)

// Commands handle special bot commands
func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if !update.Message.IsCommand() {
		return
	}

	var msg tgbotapi.MessageConfig

	switch update.Message.Command() {
	case "ping":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "pong")
		msg.ReplyToMessageID = update.Message.MessageID
	}

	bot.Send(msg)
}
