package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
)

func MatchMoveSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandMoveSnack())

	if pattern == true {
		GoToFunc = MoveSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func MoveSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(userName)
	snack = helper.CheckEmptyDayUsername(textMsg)

	if admin {
		database.MoveSnack(snack)

		if snack != "" {
			contentMessage = message.MoveSnack(snack, userName)
		} else {
			contentMessage = message.EmptyDayUsername(userName, baseCommand)
		}
	} else {
		contentMessage = message.UserAdmin(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "Cihuy") {
		sendTo = sendToGroup
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
