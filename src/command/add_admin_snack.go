package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/db"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
)

func MatchAddAdminSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandAddAdminSnack())

	if pattern == true {
		GoToFunc = AddAdminSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func AddAdminSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(userName)
	snack = helper.CheckEmptyUsername(textMsg)

	if admin {
		db.AddAdminSnack(textMsg)
	}

	if snack != "" {
		contentMessage = message.AddAdminSnack(snack, userName)
	} else {
		contentMessage = message.EmptyUsername(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "Cihuy") {
		sendTo = sendToGroup
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
