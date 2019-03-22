package main

import (
	"os"
	"time"

	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	env := godotenv.Load()
	helper.ErrorMessage(env)
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TOKEN_BOOKING"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	helper.ErrorMessage(err)

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Start()
}
