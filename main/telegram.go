package main

import (
"time"
"log"

tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "672292993:AAEY5S2ETZcc1_tCUMEPqE4GnshRwtLp3PM",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	//	websites := setupWebsitesStruct()
	scrape_website("https://www.dealabs.com/hot?page=1")

	chat := tb.Chat{ID: 726888126}
	b.Send(&chat, "je")
	b.Send(&chat, "suis")
	b.Send(&chat, "un")
	b.Send(&chat, "test")
	b.Send(&chat, "je")
	b.Send(&chat, "je")
	b.Send(&chat, "je")
	b.Send(&chat, "je")
	b.Send(&chat, "je")

	b.Start()
}
