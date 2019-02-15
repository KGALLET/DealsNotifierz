package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	PATH_TO_ALERT_FILE string = "alerts.file"
)

func displayArticle(article article) {
	fmt.Println(article.id + " | " + article.title + " | " + article.url + " | " + article.temperature)
}

// TODO define the good path
func addAlertToFile(alert []byte) {
	f, err := os.OpenFile(PATH_TO_ALERT_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	f,err := ioutil.ReadFile(PATH_TO_ALERT_FILE)
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
	fmt.Println(output)
	err = ioutil.WriteFile(PATH_TO_ALERT_FILE, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}