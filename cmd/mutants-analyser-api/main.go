package main

import (
	config2 "github.com/AgustinIzaguirre/mutants-analyser-api/internal/platform/config"
	"log"
)

func main() {
	server := SetUpServer()
	config, err := config2.GetConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := server.Run(":" + config.Port); err != nil {
		log.Fatal(err.Error())
	}
}
