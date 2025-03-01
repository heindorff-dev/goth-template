package handlers

import (
	"goth-template/server/repository"
	home "goth-template/view/pages"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	userRepository repository.UserRepository
}

func NewHomeHandler(userRepository repository.UserRepository) *HomeHandler {
	homeHandler := &HomeHandler{userRepository: userRepository}
	return homeHandler
}

func (h HomeHandler) HandleShowHome(c echo.Context) error {
	h.userRepository.Test(c.Request().Context())
	return render(c, home.Show())
}
