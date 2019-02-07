package command

import (
	"db"
	"helper"
	"message"
	"strings"
)

func MatchDoneStaging() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandDoneStaging())

	if pattern == true {
		GoToFunc = DoneStaging
	} else {
		return send_message
	}

	return GoToFunc()
}

func DoneStaging() string {
	var staging string

	staging = helper.CheckEmptyStaging(text_msg)

	db.DoneStaging((strings.ToUpper(staging)), first_name, user_id)

	if staging != "" {
		send_message = message.DoneStaging(first_name, staging)
	} else {
		send_message = message.EmptyBookDoneStaging(first_name)
	}

	return send_message
}
