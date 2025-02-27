package server

import (
	"fmt"
	"goth-template/server/helper"
	"goth-template/server/routes"

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

	routes.InitRoutes(e)

	fmt.Println("Starting server. Listening on port: " + serverPort)

	return e.Start(serverPort)
}
