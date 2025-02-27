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

	err = dbConn.Init()
	if err != nil {
		panic(err)
	}

	server := server.NewServer()
	log.Fatal(server.Start())
}
