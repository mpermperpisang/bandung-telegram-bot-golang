package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
)

func MatchHolidaySnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandHolidaySnack())

	if pattern == true {
		GoToFunc = HolidaySnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func HolidaySnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(userName)
	snack = helper.CheckEmptyUsername(textMsg)

	if admin {
		database.HolidaySnack(textMsg)

		if snack != "" {
			contentMessage = message.HolidaySnack(snack, userName)
		} else {
			contentMessage = message.EmptyUsername(userName, baseCommand)
		}
	} else {
		contentMessage = message.UserAdmin(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "diliburin") {
		sendTo = sendToGroup
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
