package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
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
	classPicture 		string
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
	picture			string
	temperature 	int
	url 			string
}


func setupWebsitesStruct() []website {
	s := []website{}

	s = append(s, website{"Dealabs", "https://www.dealabs.com/hot?page=1", "https://www.dealabs.com/nouveaux?page=1", cssClass{}})
	s = append(s, website{"Hot UK Dealz", "https://www.hotukdeals.com/hot?page=1", "https://www.hotukdeals.com/new?page=1", cssClass{}})
	s = append(s, website{"MyDealz", "https://www.mydealz.de/hot?page=1", "https://www.mydealz.de/new?page=1", cssClass{}})
	//s = append(s, website{"Pepper NL", "https://nl.pepper.com/?page=1", "https://nl.pepper.com/nieuw?page=1", cssClass{}})
	//s = append(s, website{"Chollometro", "https://www.chollometro.com/populares?page=1", "https://www.chollometro.com/nuevos?page=1", cssClass{}})

	return s
}

// TODO change to use the website struct
// TODO use emoji
func scrape_website(website website, hot bool) []article {
	soup.SetDebug(true)
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
		fmt.Print("Aucun article trouvé. Changement de la page ?")
	} else {
		fmt.Printf("Found %d articles.\n", articlesLen)
	}

	articlesToSend := []article{}
	for _, item := range articles {
		articleId := item.Attrs()["id"]
		if regexThread.MatchString(articleId) {
			// TODO class doesn't work for pepper and chollometro
			// TODO Adjust it by creating a map who will be able to give the good class attributes for a specific site
			// TODO struct
			//infoHeader := item.Find("div", "class", "threadGrid-headerMeta").FindStrict("div", "class", "flex")

			// TODO : scrap temperature

			// TODO : scrap picture

			infoTitle := item.Find("div", "class", "threadGrid-title").Find("strong", "class", "thread-title")
			a := infoTitle.Find("a")
			title := strings.TrimSpace(a.Text());
			url := strings.TrimSpace(a.Attrs()["href"])
			createdArticle := article{articleId, title, "", 1, url}
			articlesToSend = append(articlesToSend, createdArticle)
		}
	}
	return articlesToSend
}

func scrape_wanted() {
	// TODO We create a file in a specific folder which will store every keywords we want as alerts
	// TODO and then we will use the scraping function we the keywords associated in this file on the title
}


