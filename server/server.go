package server

import (
	"fmt"
	"goth-template/server/helper"
	"goth-template/server/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func getBanner() string {
	b, err := os.ReadFile("./view/public/banner.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (s *Server) Start() error {

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	serverPort, err := helper.GetEnv("SERVER_PORT")
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()

	routes.InitRoutes(e)

	fmt.Println(getBanner())

	return e.Start(serverPort)
}
