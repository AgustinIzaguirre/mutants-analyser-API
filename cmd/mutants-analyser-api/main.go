package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/controller"
)

func main() {
	server := gin.Default()
	mutantGroup := server.Group("/mutant")
	{
		mutantController := controller.New()
		mutantGroup.GET("/", mutantController.AnalyseDNA)
	}
	if error := server.Run(":5000"); error != nil {
		log.Fatal(error.Error())
	}
}