package command

var text_msg, first_name, last_name, bot_name string
var pattern, type_chat, send_message, base_command string
var user_name, chat_title string
var user_id int
var GoToFunc Func

type Func func() string

func Actions(text, username, first, last, title, BotName, command string, id int) string {
	text_msg = text
	user_name = username
	first_name = first
	last_name = last
	chat_title = title
	user_id = id
	bot_name = BotName
	base_command = command

	Group()
	Private()

	return send_message
}
