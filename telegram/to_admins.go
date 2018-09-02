package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
)

var admins = []int64{
	57772277, // @tatyshev
	33020029, // @sudoguy
}

// ToAdmins sends message to all admins
func ToAdmins(bot *tgbotapi.BotAPI, message string) {
	for _, chatID := range admins {
		msg := tgbotapi.NewMessage(chatID, message)
		bot.Send(msg)
	}
}
