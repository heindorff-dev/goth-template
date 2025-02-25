package main

import (
	"goth-template/server"
	"log"
)

func main() {
	server := server.NewServer()
	log.Fatal(server.Start())
}
