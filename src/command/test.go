package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: "528564679:AAFGFO5Aco4N1DQNRmZCJYCj0V_apSXhC7c",
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		URL:    "https://api.telegram.org",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world..aku bot yang pakai go-lang")
	})

	b.Start()
}
