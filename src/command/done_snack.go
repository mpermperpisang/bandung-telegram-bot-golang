package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/db"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchDoneSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandDoneSnack())

	if pattern == true {
		GoToFunc = DoneSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func DoneSnack() string {
	db.DoneSnack(userName, userID)

	contentMessage = message.DoneSnack(userName, firstName)
	if strings.Contains(contentMessage, "Yeay") {
		sendTo = sendToGroup
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
