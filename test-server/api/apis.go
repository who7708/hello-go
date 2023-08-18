package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func HelloGolang(c echo.Context) error {
	return c.JSON(http.StatusOK, "HelloGolang")
}
