package server

import (
	"fmt"
	"goth-template/database"
	"goth-template/server/handlers"
	"goth-template/server/repository"
	"goth-template/view"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Server struct {
	connectionManager *database.ConnectionManager
	configurations    Configurations
}

type Configurations struct {
	port string
}

func NewConfiguration(port string) *Configurations {
	config := &Configurations{
		port: port,
	}
	return config
}

func NewServer(config *Configurations, connMngr *database.ConnectionManager) *Server {
	server := &Server{
		configurations:    *config,
		connectionManager: connMngr,
	}
	return server
}

func getBanner() string {
	b, err := os.ReadFile("./view/public/banner.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (s *Server) Start() error {

	e := echo.New()
	userRepository := repository.NewUserRepository(s.connectionManager.NewClient())

	assetHandler := http.FileServer(view.GetPublicAssetsFileSystem())
	e.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", assetHandler)))

	homeHandler := handlers.NewHomeHandler(*userRepository)
	e.GET("/", homeHandler.HandleShowHome)

	fmt.Println(getBanner())

	return e.Start(s.configurations.port)
}
