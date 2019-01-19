package private

import (
	"command"
	"helper"
	"os"
)

var BotName string

func init() {
	BotName = os.Getenv("BOT_BOOKING")
}

func SendMessage(text, first, last string) string {
	var send_message string

	switch text {
	case "/start":
		send_message = command.Start(first, last)
	case "/help":
		send_message = command.Help(BotName)
	default:
		send_message = helper.Command()
	}

	return send_message
}
