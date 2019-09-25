package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	PATH_TO_WANTED_ARTICLES_FILE string = "wanted.file"
)

func displayArticle(article article, wanted bool) {
	if wanted {
		fmt.Println("WANTED " + article.id + " | " + article.title + " | " + article.url + " | " + article.temperature)
	} else {
		fmt.Println(article.id + " | " + article.title + " | " + article.url + " | " + article.temperature)
	}

}

// TODO define the good path
// TODO use it
func addAlertToFile(alert []byte) {
	f, err := os.OpenFile(PATH_TO_WANTED_ARTICLES_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(alert); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func removeAlertFromFile(alerts []string) {
	f,err := ioutil.ReadFile(PATH_TO_WANTED_ARTICLES_FILE)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(f), "\n")

	for _, alert := range alerts {
		for i, line := range lines {
			if strings.Contains(line, alert) {
				lines[i] = ""
			}
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(PATH_TO_WANTED_ARTICLES_FILE, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func readArticlesFromFile() (lines []string, err error) {
	data, err := ioutil.ReadFile(PATH_TO_WANTED_ARTICLES_FILE)
	if err != nil {
		log.Fatal(err)
		return
	}

	return strings.Split(string(data), "\n"), err
}