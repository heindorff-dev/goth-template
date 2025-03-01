package server

import (
	"fmt"
	"goth-template/database"
	"goth-template/server/handlers"
	"goth-template/server/repository"
	"goth-template/view"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Server struct {
	connectionManager *database.ConnectionManager
	configurations    Configurations
	logger            *slog.Logger
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

func New(config *Configurations, connMngr *database.ConnectionManager, logger *slog.Logger) *Server {
	server := &Server{
		configurations:    *config,
		connectionManager: connMngr,
		logger:            logger,
	}
	return server
}

func (s *Server) Start() error {
	e := echo.New()

	// Create repositories
	userRepository := repository.NewUserRepository(s.connectionManager.NewClient())

	// Public directory
	assetHandler := http.FileServer(view.GetPublicAssetsFileSystem())
	e.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", assetHandler)))

	// Set routes
	homeHandler := handlers.NewHomeHandler(*userRepository)
	e.GET("/", homeHandler.HandleShowHome)

	// Start server
	fmt.Println(getBanner())
	return e.Start(s.configurations.port)
}

func getBanner() string {
	b, err := os.ReadFile("./view/public/banner.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}
