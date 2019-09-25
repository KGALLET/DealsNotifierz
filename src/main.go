package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func MinuteTicker() *time.Ticker {
	c := make(chan time.Time, 30)
	t := &time.Ticker{C: c}
	go func() {
		for {
			n := time.Now()
			if n.Second() == 0 {
				c <- n
			}
			time.Sleep(time.Second)
		}
	}()
	return t
}

func setupWebsitesStruct() []website {
	websites := []website{}
	pepperChollometroCSS := cssClass{};
	pepperChollometroCSS.classTitle = "threadCardLayout--row--medium"
	pepperChollometroCSS.classTemperature = "threadCardLayout--row--large"

	dealabsHotUKMyDealz := cssClass{};
	dealabsHotUKMyDealz.classTitle = "threadGrid-title"
	dealabsHotUKMyDealz.classTemperature = "threadGrid-headerMeta"

	websites = append(websites, website{"Dealabs", "https://www.dealabs.com/hot?page=1", "https://www.dealabs.com/nouveaux?page=1", dealabsHotUKMyDealz})
	websites = append(websites, website{"Hot UK Dealz", "https://www.hotukdeals.com/hot?page=1", "https://www.hotukdeals.com/new?page=1", dealabsHotUKMyDealz})
	websites = append(websites, website{"MyDealz", "https://www.mydealz.de/hot?page=1", "https://www.mydealz.de/new?page=1", dealabsHotUKMyDealz})
	websites = append(websites, website{"Pepper NL", "https://nl.pepper.com/?page=1", "https://nl.pepper.com/nieuw?page=1", pepperChollometroCSS})
	websites = append(websites, website{"Chollometro", "https://www.chollometro.com/populares?page=1", "https://www.chollometro.com/nuevos?page=1", pepperChollometroCSS})

	return websites
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

	for n := range MinuteTicker().C {
		fmt.Println("Starting scrapping wanted at : ", n)
		scrapeByWanted(websites, b, &chat)
		fmt.Println("Ending scrapping wanted at : ", n)
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

	b.Start()
}
