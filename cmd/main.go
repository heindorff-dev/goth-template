package main

import (
	"goth-template/database"
	"goth-template/server"
	"log"
)

func main() {
	dbConn, err := database.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	log.Fatal(dbConn.Init())

	server := server.NewServer()
	log.Fatal(server.Start())
}
