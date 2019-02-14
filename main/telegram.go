package main

import (
	"log"
	"regexp"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)


// TODO handler add alert
// TODO handler remove alert

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

func handleAddAlert(m *tb.Message) {
	regexAdd, _ := regexp.Compile("/add ");
	alertsToAdd := regexAdd.ReplaceAllString(m.Text, "")
	alerts := strings.Split(alertsToAdd, " ")
	for _, alert := range alerts {
		addAlertToFile([]byte(alert + "\n"))
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

	b.Handle("/add", handleAddAlert)

	b.Start()
}
