package main

import (
	"database/sql"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/controller"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/persistence"
	config2 "github.com/AgustinIzaguirre/mutants-analyser-api/internal/platform/config"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func SetUpServer() *gin.Engine {
	server := gin.Default()
	mutantGroup := server.Group("/mutant")
	{
		// TODO make table name constant
		mutantController := controller.New(persistence.New("analysis", PostgresConnectionProvider))
		mutantGroup.GET("/", mutantController.AnalyseDNA)
	}
	return server
}

func PostgresConnectionProvider() (*sql.DB, error) {
	config, err := config2.GetConfig(".")
	if err != nil {
		log.Fatal(err)
		return &sql.DB{}, err
	}
	databaseSourceName := "postgres://" + config.PostgresUsername + ":" + config.PostgresPassword +
							"@" + config.PostgresServerAddress + ":" + config.PostgresServerPort + "/" +
							config.PostgresDatabase + "?sslmode=" + config.PostgresSslMode
	database, err := sql.Open("postgres", databaseSourceName)
	if err != nil {
		log.Fatal(err)
		return &sql.DB{}, err
	}
	return database, nil
}