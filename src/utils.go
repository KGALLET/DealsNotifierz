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

func displayArticle(article article) {
	fmt.Println(article.id + " | " + article.title + " | " + article.url + " | " + article.temperature)
}

// TODO define the good path
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

// TODO define the good path
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

	// TODO remove trailling spaces inside file (not beginning and end of file)
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(PATH_TO_WANTED_ARTICLES_FILE, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func readArticlesFromFile() {
	// TODO read file with the wanted articles keyword and return it as an array
	f,err := ioutil.ReadFile(PATH_TO_WANTED_ARTICLES_FILE)
	if err != nil {
		log.Fatal(err)
	}

	str := string(f)

	fmt.Println(str)

}