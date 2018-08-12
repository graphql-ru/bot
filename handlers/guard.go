package handlers

import (
	"gopkg.in/telegram-bot-api.v4"
)

// Guard restricts bot usage only in private chats and @graphql_ru
func Guard(bot *tgbotapi.BotAPI, update tgbotapi.Update, next func()) {
	chat := update.Message.Chat

	if !chat.IsGroup() || !chat.IsSuperGroup() || !chat.IsChannel() {
		next()
	}

	if chat.UserName == "graphql_ru" {
		next()
	}
}
