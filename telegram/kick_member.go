package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
)

// KickMember restrict chat user
func KickMember(bot *tgbotapi.BotAPI, chatID int64, userID int) {
	memberConfig := tgbotapi.ChatMemberConfig{
		ChatID: chatID,
		UserID: userID,
	}

	config := tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: memberConfig,
	}

	bot.KickChatMember(config)
}
