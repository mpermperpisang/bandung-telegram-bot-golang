package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchDoneStaging() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandDoneStaging())

	if pattern {
		GoToFunc = DoneStaging
	} else {
		return "not match"
	}

	return GoToFunc()
}

func DoneStaging() string {
	var staging string

	staging = helper.CheckEmptyStaging(textMsg)

	database.DoneStaging((strings.ToUpper(staging)), firstName, userID)

	if staging != "" {
		contentMessage = message.DoneStaging(userName, strings.Trim(staging, " "))
	} else {
		contentMessage = message.EmptyBookDoneStaging(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "selesai") {
		sendTo = msg.Chat
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
