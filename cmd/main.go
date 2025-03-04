package main

import (
	"goth-template/database"
	"goth-template/internal/helper"
	"goth-template/internal/server"
	"log"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	databaseConnectionCredentials := mustGetDatabaseCredentials()
	databaseConnectionConfiguration := mustGetDatabaseConfiguration()
	databaseConnectionManager, err := database.NewConnectionManager(*databaseConnectionConfiguration, *databaseConnectionCredentials, logger)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	pingRes := databaseConnectionManager.Ping()
	if pingRes != nil {
		logger.Error(pingRes.Error())
		panic(pingRes)
	}

	serverConfiguration := mustGetServerConfiguration()
	server := server.New(serverConfiguration, databaseConnectionManager, logger)

	log.Fatal(server.Start())
}

func mustGetDatabaseCredentials() *database.ConnectionCredential {
	credUsername := helper.MustGetEnv("DB_USERNAME")
	credPassword := helper.MustGetEnv("DB_PASSWORD")
	creds := database.NewConnectionCredentials(credUsername, credPassword)
	return creds
}

func mustGetDatabaseConfiguration() *database.ConnectionConfiguration {
	dbName := helper.MustGetEnv("DB_DATABASE")
	host := helper.MustGetEnv("DB_CONTAINER")
	port := helper.MustGetEnv("DB_REMOTE_PORT")
	protocol := helper.MustGetEnv("DB_PROTOCOL")
	config := database.NewConnectionConfiguration(dbName, host, port, protocol)
	return config
}

func mustGetServerConfiguration() *server.Configurations {
	serverPort := ":" + helper.MustGetEnv("API_REMOTE_PORT")
	config := server.NewConfiguration(serverPort)
	return config
}
