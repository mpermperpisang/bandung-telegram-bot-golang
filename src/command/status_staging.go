package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/db"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchStatusStaging() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandStatusStaging())

	if pattern == true {
		GoToFunc = StatusStaging
	} else {
		return "not match"
	}

	return GoToFunc()
}

func StatusStaging() string {
	var staging string

	staging = helper.CheckEmptyStaging(textMsg)

	db.StatusStaging(strings.ToUpper(staging))

	if staging != "" {
		contentMessage = message.StagingStatus(strings.Trim(staging, " "))
	} else {
		contentMessage = message.EmptyStatusStaging(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "Status") {
		sendTo = sendToGroup
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
