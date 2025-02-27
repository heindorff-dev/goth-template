package main

import (
	"goth-template/database"
	"goth-template/server"
	"log"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	db.Begin() //idk?

	server := server.NewServer()
	log.Fatal(server.Start())
}
