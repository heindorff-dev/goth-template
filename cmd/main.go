package main

import (
	"goth-template/database"
	"goth-template/server"
	"goth-template/server/helper"
	"log"
)

func main() {
	databaseConnectionCredentials := mustGetDatabaseCredentials()
	databaseConnectionConfiguration := mustGetDatabaseConfiguration()
	databaseConnectionManager, err := database.NewConnectionManager(*databaseConnectionConfiguration, *databaseConnectionCredentials)

	if err != nil {
		panic("Failed to connect to database")
	}

	serverConfiguration := mustGetServerConfiguration()
	server := server.NewServer(serverConfiguration, databaseConnectionManager)

	log.Fatal(server.Start())
}

func mustGetDatabaseCredentials() *database.ConnectionCredentials {
	credUsername := helper.MustGetEnv("DB_USERNAME")
	credPassword := helper.MustGetEnv("DB_PASSWORD")
	creds := database.NewConnectionCredentials(credUsername, credPassword)
	return creds
}

func mustGetDatabaseConfiguration() *database.ConnectionConfigurations {
	dbName := helper.MustGetEnv("DB_DATABASE")
	host := helper.MustGetEnv("DB_CONTAINER")
	port := helper.MustGetEnv("DB_REMOTE_PORT")
	config := database.NewConnectionConfiguration(dbName, host, port)
	return config
}

func mustGetServerConfiguration() *server.Configurations {
	serverPort := ":" + helper.MustGetEnv("API_REMOTE_PORT")
	config := server.NewConfiguration(serverPort)
	return config
}
