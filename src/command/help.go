package command

import (
	"helper"
	"message"
	"os"
	"strings"
)

func MatchHelp() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandHelp())

	if pattern == true {
		GoToFunc = Help
	} else {
		return send_message
	}

	return GoToFunc()
}

func Help() string {
	var content string

	switch bot_name {
	case os.Getenv("BOT_BOOKING"):
		content = message.Booking()
	default:
		content = helper.Help()
	}

	send_message = message.Header() + content + message.Footer()

	return send_message
}
