package command

import (
	"helper"
)

var text_msg, first_name, last_name, bot_name string
var pattern, type_chat, send_message string
var user_id int
var GoToFunc Func

type Func func() string

func Actions(text, first, last, BotName string, id int) string {
	text_msg = text
	first_name = first
	last_name = last
	user_id = id
	bot_name = BotName

	Group()
	Private()

	return send_message
}

func NotFound() string {
	send_message = helper.Command()

	return send_message
}
