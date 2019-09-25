package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"regexp"
	"strings"
)

func scrapeByWanted(websites []website, bot *tb.Bot, recipient tb.Recipient) {
	fmt.Println("Scrapping for wanted items")
	wantedArticles := []article{}
	for _, website := range websites {
		scrappedArticles := scrape_wanted(website)
		if (len(scrappedArticles) != 0) {
			wantedArticles = append(wantedArticles, scrappedArticles...)
		}
	}

	for _, wantedArticle := range wantedArticles {
		displayArticle(wantedArticle, true)
	}

	sendMessageByArticle(bot, recipient, wantedArticles)
	bot.Send(recipient, "New(s) wanted articles scrapped :D")
}

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
		bot.Send(recipient, article.title + " | " + "URL : " + article.url)
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

func handleRemoveAlert(m *tb.Message) {
	regexAdd, _ := regexp.Compile("/del ");
	alertsToRemove := regexAdd.ReplaceAllString(m.Text, "")
	alerts := strings.Split(alertsToRemove, " ")
	removeAlertFromFile(alerts)
}

