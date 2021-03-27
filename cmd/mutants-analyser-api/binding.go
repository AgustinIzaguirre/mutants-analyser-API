package main

import (
	"database/sql"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/controller"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/persistence"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/service"
	controller2 "github.com/AgustinIzaguirre/mutants-analyser-api/internal/stats/controller"
	config2 "github.com/AgustinIzaguirre/mutants-analyser-api/internal/platform/config"
	persistence2 "github.com/AgustinIzaguirre/mutants-analyser-api/internal/stats/persistence"
	service2 "github.com/AgustinIzaguirre/mutants-analyser-api/internal/stats/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

const ANALYSIS_TABLE_NAME string = "analysis"

func SetUpServer() *gin.Engine {
	server := gin.Default()
	mutantGroup := server.Group("/mutant")
	{
		mutantController := controller.New(
								service.New(
									persistence.New(ANALYSIS_TABLE_NAME, PostgresConnectionProvider)))
		mutantGroup.POST("/", mutantController.AnalyseDNA)
	}

	StatsGroup := server.Group("/stats")
	{
		statsController := controller2.New(
			service2.New(
				persistence2.New(ANALYSIS_TABLE_NAME, PostgresConnectionProvider)))
		StatsGroup.GET("/", statsController.GetStats)
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