package command

import (
	"helper"
	"message"
	"strings"
)

func MatchStart() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandStart())

	if pattern == true {
		GoToFunc = Start
	} else {
		return send_message
	}

	return GoToFunc()
}

func Start() string {
	send_message = message.Start(first_name, last_name)

	return send_message
}
