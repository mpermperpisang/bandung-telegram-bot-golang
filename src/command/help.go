package command

import (
	"helper"
	"message"
	"os"
)

var BotBooking string

func init() {
	BotBooking = os.Getenv("BOT_BOOKING")
}

func Help(name string) string {
	var content string

	switch name {
	case BotBooking:
		content = message.Booking()
	default:
		content = helper.Help()
	}

	var msg_help = message.Header() + content + message.Footer()

	return msg_help
}
