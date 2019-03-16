package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/db"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
)

func MatchDeleteSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandDeleteSnack())

	if pattern == true {
		GoToFunc = DeleteSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func DeleteSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(firstName)
	snack = helper.CheckEmptyUsername(textMsg)

	if admin {
		db.DeleteSnack(textMsg)
	}

	if snack != "" {
		contentMessage = message.DeleteSnack(snack, firstName)
	} else {
		contentMessage = message.EmptyUsername(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "Colek") {
		sendTo = sendToGroup
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
