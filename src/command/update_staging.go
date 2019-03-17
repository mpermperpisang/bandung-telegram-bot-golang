package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/db"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchUpdateStaging() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandUpdateStaging())

	if pattern == true {
		GoToFunc = UpdateStaging
	} else {
		return "not match"
	}

	return GoToFunc()
}

func UpdateStaging() string {
	var staging string

	staging = helper.CheckEmptySquadStaging(textMsg)

	db.UpdateStaging(strings.ToUpper(staging))

	if staging != "" {
		contentMessage = message.UpdateStaging(strings.Trim(staging, " "))
	} else {
		contentMessage = message.EmptySquadStaging(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "berhasil") {
		sendTo = msg.Chat
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
