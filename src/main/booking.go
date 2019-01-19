package main

import (
	"chat_type/group"
	"chat_type/private"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

var token string

func init() {
	env := godotenv.Load()

	if env != nil {
		log.Fatal("Error loading .env file")
	}

	token = os.Getenv("TOKEN_BOOKING")
}

func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle(tb.OnText, func(m *tb.Message) {
		if !m.Private() {
			bot.Send(m.Chat, group.SendMessage(m.Text, m.Sender.Username), tb.ModeHTML)
		} else {
			bot.Send(m.Sender, private.SendMessage(m.Text, m.Chat.FirstName, m.Chat.LastName), tb.ModeHTML)
		}
	})

	bot.Start()
}
