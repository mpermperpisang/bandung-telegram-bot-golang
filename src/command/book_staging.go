package command

import (
	"db"
	"helper"
	"message"
	"strings"
)

func MatchBookStaging() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandBookStaging())

	if pattern == true {
		GoToFunc = BookStaging
	} else {
		return send_message
	}

	return GoToFunc()
}

func BookStaging() string {
	var staging string

	staging = helper.CheckEmptyStaging(text_msg)

	db.BookStaging((strings.ToUpper(staging)), first_name, user_id)

	if staging != "" {
		send_message = message.BookStaging(first_name, staging)
	} else {
		send_message = message.EmptyBookDoneStaging(first_name)
	}

	return send_message
}
