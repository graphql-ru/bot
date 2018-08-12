package handlers

import (
	"log"

	"github.com/graphql-ru/bot/telegram"
	"gopkg.in/telegram-bot-api.v4"
)

// Join handle joined members
func Join(bot *tgbotapi.BotAPI, update tgbotapi.Update, next func()) {
	var members []tgbotapi.User

	if update.Message.NewChatMembers == nil {
		next()
		return
	}

	members = *update.Message.NewChatMembers

	for _, member := range members {
		if telegram.IsChinaBot(&member) {
			log.Printf("[KICK] @%s", member.UserName)

			telegram.DeleteMessage(bot, update.Message)
			telegram.KickMember(bot, update.Message.Chat.ID, member.ID)
		}
	}

	next()
}
