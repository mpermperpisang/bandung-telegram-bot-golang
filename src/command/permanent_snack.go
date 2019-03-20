package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
)

func MatchPermanentSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandPermanentSnack())

	if pattern == true {
		GoToFunc = PermanentSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func PermanentSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(userName)
	snack = helper.CheckEmptyDayUsername(textMsg)

	if admin {
		database.PermanentSnack(snack)

		if snack != "" {
			contentMessage = message.PermanentSnack(snack, userName)
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
