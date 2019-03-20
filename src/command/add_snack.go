package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
)

func MatchAddSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandAddSnack())

	if pattern == true {
		GoToFunc = AddSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func AddSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(userName)
	snack = helper.CheckEmptyDayUsername(textMsg)

	if admin {
		database.AddSnack(snack)
	}

	if snack != "" {
		contentMessage = message.AddSnack(snack, userName)
	} else {
		contentMessage = message.EmptyDayUsername(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "Cihuy") {
		sendTo = sendToGroup
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
