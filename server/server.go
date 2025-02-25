package server

import (
	"fmt"
	"goth-template/server/handlers"
	"goth-template/server/helper"
	"goth-template/view"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	serverPort, err := helper.MustGetEnv("SERVER_PORT")
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()

	// Handle static assets
	assetHandler := http.FileServer(view.GetPublicAssetsFileSystem())
	e.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", assetHandler)))

	homeHandler := handlers.HomeHandler{}
	e.GET("/", homeHandler.HandleShowHome)

	fmt.Println("Starting server. Listening on port: " + serverPort)

	return e.Start(serverPort)
}
