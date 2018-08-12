package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
)

// DeleteMessage deletes message and restrict author
func DeleteMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	deleteConfig := tgbotapi.DeleteMessageConfig{
		ChatID:    message.Chat.ID,
		MessageID: message.MessageID,
	}

	bot.DeleteMessage(deleteConfig)
}
