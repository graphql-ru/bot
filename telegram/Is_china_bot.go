package telegram

import (
	"regexp"

	"gopkg.in/telegram-bot-api.v4"
)

// IsChinaBot check user name looks like china bot
func IsChinaBot(user *tgbotapi.User) bool {
	re := regexp.MustCompile("VX,QQ")
	name := user.FirstName

	if len(name) > 200 {
		return true
	}

	if re.MatchString(name) {
		return true
	}

	return false
}
