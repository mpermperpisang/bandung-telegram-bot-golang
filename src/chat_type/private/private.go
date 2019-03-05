package private

import "command"

func SendMessage(text, username, first, last, title, BotName, baseCommand string, id int) string {
	var send_message string

	send_message = command.Actions(text, username, first, last, title, BotName, baseCommand, id)

	return send_message
}
