package helper

import (
	"message"
)

func Help() string {
	return message.DefaultHelp()
}

func Command() string {
	return message.DefaultCommand()
}
