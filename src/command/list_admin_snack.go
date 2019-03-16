package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/db"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchListAdminSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandListAdminSnack())

	if pattern == true {
		GoToFunc = ListAdminSnack
	} else {
		return contentMessage
	}

	return GoToFunc()
}

func ListAdminSnack() string {
	db.ListAdminSnack()

	contentMessage = message.ListAdminSnack(userName)
	sendTo = sendToGroup

	return "success"
}
