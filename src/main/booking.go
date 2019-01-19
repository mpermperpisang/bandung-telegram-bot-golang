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

func main() {
	env := godotenv.Load()

	if env != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN_BOOKING")
	name := os.Getenv("BOT_BOOKING")

	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle(tb.OnText, func(m *tb.Message) {
		if m.Private() {
			bot.Send(m.Sender, private.SendMessage(name, m.Text, m.Chat.FirstName, m.Chat.LastName), tb.ModeHTML)
		} else {
			group.SendMessage(m.Text)
		}
	})

	bot.Start()
}
