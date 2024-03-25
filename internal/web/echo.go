package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Echo() error {
	renderer, err := NewRenderer()
	if err != nil {
		return err
	}
	e := echo.New()
	e.Renderer = renderer
	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "pages/index", map[string]string{
			"PageTitle": "Home",
		})
	})
	e.GET("/search", SearchGet)
	e.GET("/dashboard", DashboardGet)

	e.Logger.Fatal(e.Start(":8080"))
	return nil
}
