package handler

import (
	"goth-template/view/home"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleShowHome(c echo.Context) error {
	return render(c, home.Show())
}
