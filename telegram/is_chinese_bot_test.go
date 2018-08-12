package telegram_test

import (
	"testing"

	"github.com/graphql-ru/bot/telegram"
	"github.com/stretchr/testify/assert"
	"gopkg.in/telegram-bot-api.v4"
)

func TestIsChineseBot(t *testing.T) {
	var member *tgbotapi.User

	member = memberWithfirstName("╋VX,QQ（同号）：253239090 专业工作室推广拉人【电报群拉国内外有无username都可拉、指定群拉人】【机器人定制】【社群代运营】【twitter关注、转发】【facebook关注、转发】【youtube点赞、评论】【出售成品电报账号】 （欢迎社群运营者、项目方、交易所洽谈合作）优质空投分享QQ群473157472 本工作室全网最低价、服务最好、活人质量最高 招收代理")
	assert.Equal(t, true, telegram.IsChineseBot(member))

	member = memberWithfirstName("VX,QQ（同号）：253239090 专业工作室推广拉人【电报群国内外有无username、指定群拉人】【机器人定制】【社群代运营】【twitter关注、转发】【facebook关注、转发】【youtube点赞、评论】【出售成品电报账号】 （欢迎社群运营者、项目方、交易所洽谈合作）本工作室全网最低价、服务最好、活人质量最高 招收代理 Telegram with username:smartworkshop or Email: smartelegram at outlook.com.Get REAL ACTIVE HUMAN members to promote your TELEGRAM GROUP or CHANNEL. Fabulous services and the BEST -60% DISCOUNT.(Please tell if this msg bothers)")
	assert.Equal(t, true, telegram.IsChineseBot(member))

	member = memberWithfirstName("专业工作室推广拉人【电报群国内外有无username、指定群拉人】【机器人定制】【社群代运营】【twitter关注、转发】【facebook关注、转发】【youtube点赞、评论】【出售成品电报账号】 （欢迎社群运营者、项目方、交易所洽谈合作）本工作室全网最低价、服务最好、活人质量最高 招收代理")
	assert.Equal(t, true, telegram.IsChineseBot(member))

	member = memberWithfirstName("VX,QQ 专业工作室推广拉人【电报群国内外有无username、指定群拉人")
	assert.Equal(t, true, telegram.IsChineseBot(member))

	member = memberWithfirstName("专业工作室推广拉人")
	assert.Equal(t, false, telegram.IsChineseBot(member))

	member = memberWithfirstName("foo")
	assert.Equal(t, false, telegram.IsChineseBot(member))
}

func memberWithfirstName(name string) *tgbotapi.User {
	return &tgbotapi.User{
		FirstName: name,
	}
}
