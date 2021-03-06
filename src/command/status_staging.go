package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchStatusStaging() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandStatusStaging())

	if pattern {
		GoToFunc = StatusStaging
	} else {
		return "not match"
	}

	return GoToFunc()
}

func StatusStaging() string {
	var staging string

	staging = helper.CheckEmptyStaging(textMsg)

	database.StatusStaging(strings.ToUpper(staging))

	if staging != "" {
		contentMessage = message.StagingStatus(strings.Trim(staging, " "))
	} else {
		contentMessage = message.EmptyStatusStaging(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "Status") {
		sendTo = msg.Chat
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
