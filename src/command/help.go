package command

import (
	"helper"
	"message"
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
	case "stg_book_bot":
		content = message.Booking()
	default:
		content = helper.Help()
	}

	send_message = message.Header() + content + message.Footer()

	return send_message
}
