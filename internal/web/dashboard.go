package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func DashboardGet(c echo.Context) error {
	data := map[string]interface{}{}
	return c.Render(http.StatusOK, "pages/dashboard", data)
}
