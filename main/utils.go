package main

import (
	"log"
	"os"
)

// TODO define the good path
func addAlertToFile(alert []byte) {
	f, err := os.OpenFile("alerts.file", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
