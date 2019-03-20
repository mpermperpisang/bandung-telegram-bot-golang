package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
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
	database.BaikSnack(userName)

	contentMessage = message.BaikSnack(userName)
	sendTo = msg.Chat

	return "success"
}
