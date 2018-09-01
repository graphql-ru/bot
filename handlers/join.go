package handlers

import (
	"fmt"
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

	ChatID := update.Message.Chat.ID
	UserName := update.Message.From.UserName

	log.Printf("[JOIN] [ChatID: %d] @%s", ChatID, UserName)

	members = *update.Message.NewChatMembers

	for _, member := range members {
		if telegram.IsChineseBot(&member) {
			log.Printf("[KICK] [ChatID: %d] @%s", ChatID, UserName)

			telegram.DeleteMessage(bot, update.Message)
			telegram.KickMember(bot, update.Message.Chat.ID, member.ID)
			telegram.SendToAdmins(bot, fmt.Sprintf("Я кикнул @%s", UserName))
		}
	}

	next()
}
