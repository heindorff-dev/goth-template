package main

import (
	"goth-template/server/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	homeHandler := handler.HomeHandler{}

	e.GET("/", homeHandler.HandleShowHome)

	e.Logger.Fatal(e.Start(":5000"))
}
