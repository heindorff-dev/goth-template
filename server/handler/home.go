package handler

import (
	"goth-template/server/service"
	home "goth-template/view/page"
	"log/slog"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	userService *service.UserService
	logger      slog.Logger
}

func NewHomeHandler(userService *service.UserService, logger *slog.Logger) *HomeHandler {
	homeHandler := &HomeHandler{
		userService: userService,
		logger:      *logger,
	}
	return homeHandler
}

func (h HomeHandler) HandleShowHome(c echo.Context) error {
	h.logger.Info("HomeHandler.HandleShowHome()")
	h.userService.GetUsers(c.Request().Context())
	return render(c, home.Show())
}
