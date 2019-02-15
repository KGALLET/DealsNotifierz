package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"regexp"
	"strings"
)

// TODO mise en page du message
// TODO use emoji

func scrapeByType(category string, websites []website, bot *tb.Bot, recipient tb.Recipient) {
	for _, website := range websites {
		articles := []article{}
		if category == "hot" {
			articles = scrape_website(website, true)
		}  else {
			articles = scrape_website(website, false)
		}
		fmt.Println(articles)
/*		for _, article := range articles {
			displayArticle(article)
		}*/
		/*sendMessageByArticle(bot, recipient, articles)
		bot.Send(recipient, "================================================================================" +
								  "================================================================================")*/
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

