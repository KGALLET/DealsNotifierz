package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)


// TODO handler add alert
// TODO handler remove alert
// TODO handler get instant hot deal
// TODO handler get instant news deal

func scrapeByType(category string, websites []website, bot *tb.Bot, recipient tb.Recipient) {
	for _, website := range websites {
		articles := []article{}
		if category == "hot" {
			articles = scrape_website(website, true)
		}  else {
			articles = scrape_website(website, false)
		}
		sendMessageByArticle(bot, recipient, articles)
		bot.Send(recipient, "================================================================================" +
								  "================================================================================")
	}
}

func sendMessageByArticle(bot *tb.Bot, recipient tb.Recipient, articles []article) {
	for _, article := range articles {
		bot.Send(recipient, article.title + " | " + "URL : " + article.url, tb.NoPreview)
	}
}

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


	b.Handle("/hot", func(m *tb.Message) {
		scrapeByType("hot", websites, b, &chat)
	})

	b.Handle("/new", func(m *tb.Message) {
		scrapeByType("new", websites, b, &chat)
	})

	//articles := scrape_website("https://www.dealabs.com/hot?page=1")


	//for _, article := range articles {
	//	sendMessage(b, &chat, article.title + " | " + "URL : " + article.url)
	//}

	b.Start()
}
