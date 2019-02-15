package main

import (
	"log"
	"time"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	websites := setupWebsitesStruct()

	b, err := tb.NewBot(tb.Settings{
		Token:  "672292993:AAEY5S2ETZcc1_tCUMEPqE4GnshRwtLp3PM",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	chat := tb.Chat{ID: 726888126}

	if err != nil {
		log.Fatal(err)
		return
	}

	// Handlers
	b.Handle("/hot", func(m *tb.Message) {
		scrapeByType("hot", websites, b, &chat)
	})
	b.Handle("/new", func(m *tb.Message) {
		scrapeByType("new", websites, b, &chat)
	})
	b.Handle("/add", handleAddAlert)
	b.Handle("/del", handleRemoveAlert)

	/*scrapeByType("hot", websites, b, &chat)*/
	getWantedArticles()

	b.Start()
}