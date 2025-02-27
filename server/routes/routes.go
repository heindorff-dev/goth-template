package routes

import (
	"goth-template/server/handlers"
	"goth-template/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	assetHandler := http.FileServer(view.GetPublicAssetsFileSystem())
	e.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", assetHandler)))

	homeHandler := handlers.HomeHandler{}
	e.GET("/", homeHandler.HandleShowHome)
}
