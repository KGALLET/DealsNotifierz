package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	USER_AGENT_HEADER string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/68.0.3440.106 " + "Safari/537.36 OPR/55.0.2994.61"
)

type cssClass struct {
	classTitle			string
	classTemperature	string
}

type website struct {
	title 			string
	hotUrl			string
	newUrl			string
	cssClasses		cssClass
}

type article struct {
	id				string
	title			string
	temperature 	string
	url 			string
}

func scrape_website(website website, hot bool) []article {
	soup.Headers["user-agent"] = USER_AGENT_HEADER
	url := ""
	fmt.Println("Scrapping website " + website.title)
	if hot {
		url = website.hotUrl
	} else {
		url = website.newUrl
	}

	resp, err := soup.Get(url)
	if err != nil {
		fmt.Println("Error, killing the process", err)
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)

	regexThread, _ := regexp.Compile("thread_.*");
	articles := doc.FindAll("article")
	articlesLen := len(articles)
	if articles == nil || articlesLen == 0 {
		fmt.Print("No article found. The page has changed ?")
	} else {
		fmt.Printf("Found %d articles.\n", articlesLen)
	}

	articlesToSend := []article{}
	for _, item := range articles {
		articleId := item.Attrs()["id"]
		var createdArticle *article = &article{articleId, "", "", ""}
		if regexThread.MatchString(articleId) {
			scrape_titleAndUrl(item, createdArticle, website)
			scrape_temperature(item, createdArticle, website)
			articlesToSend = append(articlesToSend, *createdArticle)
		}
	}

	return articlesToSend
}

func scrape_titleAndUrl(item soup.Root, article *article, website website) {
	infoTitle := item.Find("div", "class", website.cssClasses.classTitle).Find("strong", "class", "thread-title")

	a := infoTitle.Find("a")

	// in this case, that mean that this is not an article that we can buy,
	// but just some ads (for example on pepper.nl)
	// if we do not catch this error, than segfault
	if a.Error != nil {
		return
	}
	title := strings.TrimSpace(a.Text())
	url := strings.TrimSpace(a.Attrs()["href"])
	article.title = title
	article.url = url
}


func scrape_temperature(item soup.Root, article *article, website website) {
	infoTemperature := item.Find("div", "class", website.cssClasses.classTemperature).FindAll("span")

 	regexClassTemp, _ := regexp.Compile("cept-vote-temp vote-temp [A-Za-z\\s\\-]*");
	for _, DOMElement := range infoTemperature {
		if regexClassTemp.MatchString(DOMElement.Attrs()["class"]) {
			article.temperature = strings.TrimSpace(DOMElement.Text())
		}
	}
}


func getWantedArticles() {
//	lines, err := readArticlesFromFile()
//	if err != nil {
//		log.Fatal(err)
	//	}
	//for _, line := range lines{
	//	fmt.Println(string(line))
	//}
}

// TODO scrape by alert wanted
func scrape_wanted(alreadyFoundDeals []article, wantedArticles []article ) {
	// TODO We create a file in a specific folder which will store every keywords we want as alerts
	// TODO and then we will use the scraping function we the keywords associated in this file on the title




}


