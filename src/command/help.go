package command

import (
	"os"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchHelp() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandHelp())

	if pattern {
		GoToFunc = Help
	} else {
		return "not match"
	}

	return GoToFunc()
}

func Help() string {
	var content string

	switch botName {
	case os.Getenv("BOT_BOOKING"):
		content = message.Booking()
	case os.Getenv("BOT_SNACK"):
		content = message.Snack()
	default:
		content = helper.Help()
	}

	contentMessage = message.Header() + content + message.Footer()
	sendTo = sendToPrivate

	return "success"
}
