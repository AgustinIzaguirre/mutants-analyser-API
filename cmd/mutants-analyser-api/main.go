package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()
	if error := server.Run(":5000"); error != nil {
		log.Fatal(error.Error())
	}
}