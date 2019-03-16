package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchStart() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandStart())

	if pattern == true {
		GoToFunc = Start
	} else {
		return "failed"
	}

	return GoToFunc()
}

func Start() string {
	contentMessage = message.Start(firstName, lastName)
	sendTo = sendToPrivate

	return "success"
}
