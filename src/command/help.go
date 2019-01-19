package command

import (
	"message"
	"os"
)

func Help(name string) string {
	var content string

	switch name {
	case os.Getenv("BOT_BOOKING"):
		content = message.Booking()
	default:
		content = message.DefaultHelp()
	}

	var msg_help = message.Header() + content + message.Footer()
	return msg_help
}
