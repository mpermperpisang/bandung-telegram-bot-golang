package private

import (
	"command"
)

func SendMessage(text, first, last, BotName string, id int) string {
	var send_message string

	send_message = command.Actions(text, first, last, BotName, id)

	return send_message
}
