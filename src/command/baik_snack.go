package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
)

func MatchBaikSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandBaikSnack())

	if pattern == true {
		GoToFunc = BaikSnack
	} else {
		return contentMessage
	}

	return GoToFunc()
}

func BaikSnack() string {
	var admin bool

	admin = user.IsAdmin(userName)

	if admin {
		database.BaikSnack(userName)

		contentMessage = message.BaikSnack(userName)
	} else {
		contentMessage = message.UserAdmin(userName, baseCommand)
	}

	sendTo = msg.Chat

	return "success"
}
