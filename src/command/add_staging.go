package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchAddStaging() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandAddStaging())

	if pattern {
		GoToFunc = AddStaging
	} else {
		return "not match"
	}

	return GoToFunc()
}

func AddStaging() string {
	var staging string

	staging = helper.CheckEmptySquadStaging(textMsg)

	database.AddStaging(strings.ToUpper(staging))

	if staging != "" {
		contentMessage = message.AddStaging(strings.Trim(staging, " "))
	} else {
		contentMessage = message.EmptySquadStaging(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "Status :") {
		sendTo = msg.Chat
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
