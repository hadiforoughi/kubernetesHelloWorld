package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func getBody(c echo.Context) error  {
	return c.String(http.StatusOK, "get body")
}
