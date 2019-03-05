package command

import (
	"db"
	"helper"
	"message"
	"strings"
)

func MatchDoneSnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandDoneSnack())

	if pattern == true {
		GoToFunc = DoneSnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func DoneSnack() string {
	db.DoneSnack(user_name, user_id)

	send_message = message.DoneSnack(user_name, first_name)

	return send_message
}
