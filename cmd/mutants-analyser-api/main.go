package main

import (
	"log"
)

func main() {
	server := SetUpServer()
	if error := server.Run(":5000"); error != nil {
		log.Fatal(error.Error())
	}
}
