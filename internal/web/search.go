package web

import (
	"github.com/labstack/echo/v4"
	"github.com/vatusa/training/internal/data"
	"net/http"
)

type SearchData struct {
	MetaData
	Controllers []data.Controller
}

func SearchGet(c echo.Context) error {
	return c.Render(http.StatusOK, "pages/search", SearchData{
		MetaData:    MakeMetaData(),
		Controllers: nil,
	})
}
