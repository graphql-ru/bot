package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
)

const graphqlChatID int64 = -1001079441526 // @graphql_ru

// ToGraphQl sends message to graphql_ru chat
func ToGraphQl(bot *tgbotapi.BotAPI, message string) {
	msg := tgbotapi.NewMessage(graphqlChatID, message)
	msg.DisableWebPagePreview = true
	msg.ParseMode = "markdown"
	bot.Send(msg)
}
