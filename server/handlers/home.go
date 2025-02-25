package handlers

import (
	home "goth-template/view/pages"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleShowHome(c echo.Context) error {
	return render(c, home.Show())
}
