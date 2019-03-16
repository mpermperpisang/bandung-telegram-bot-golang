package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/db"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
)

func MatchCancelSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandCancelSnack())

	if pattern == true {
		GoToFunc = CancelSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func CancelSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(firstName)
	snack = helper.CheckEmptyUsername(textMsg)

	if admin {
		db.CancelSnack(textMsg)
	}

	if snack != "" {
		contentMessage = message.CancelSnack(snack, firstName)
	} else {
		contentMessage = message.EmptyUsername(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "dibatalin") {
		sendTo = sendToGroup
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
