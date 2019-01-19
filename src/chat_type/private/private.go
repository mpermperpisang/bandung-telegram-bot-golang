package private

import (
	"command"
	"message"
)

func SendMessage(name, text, first, last string) string {
	var value string

	switch text {
	case "/start":
		value = command.Start(first, last)
	case "/help":
		value = command.Help(name)
	default:
		value = message.DefaultCommand()
	}

	return value
}
