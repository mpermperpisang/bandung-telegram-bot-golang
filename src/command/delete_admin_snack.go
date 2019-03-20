package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
)

func MatchDeleteAdminSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandDeleteAdminSnack())

	if pattern == true {
		GoToFunc = DeleteAdminSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func DeleteAdminSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(userName)
	snack = helper.CheckEmptyUsername(textMsg)

	if admin {
		database.DeleteAdminSnack(textMsg)

		if snack != "" {
			contentMessage = message.DeleteAdminSnack(snack, userName)
		} else {
			contentMessage = message.EmptyUsername(userName, baseCommand)
		}
	} else {
		contentMessage = message.UserAdmin(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "Yaah") {
		sendTo = sendToGroup
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
