package main

import (
	"fmt"
	"goth-template/server/handler"
	"goth-template/server/helper"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	serverPort, err := helper.MustGetEnv("SERVER_PORT")
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()

	homeHandler := handler.HomeHandler{}
	e.GET("/", homeHandler.HandleShowHome)

	fmt.Println("Starting server. Listening on port: " + serverPort)
	e.Start(serverPort)
}
