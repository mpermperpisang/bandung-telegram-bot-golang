package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchListAdminSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandListAdminSnack())

	if pattern {
		GoToFunc = ListAdminSnack
	} else {
		return contentMessage
	}

	return GoToFunc()
}

func ListAdminSnack() string {
	database.ListAdminSnack()

	contentMessage = message.ListAdminSnack(userName)
	sendTo = msg.Chat

	return "success"
}
