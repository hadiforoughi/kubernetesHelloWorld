package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func router(e *echo.Echo) {
	// /home
	e.GET("/healthz", func(c echo.Context) error { return c.NoContent(http.StatusNoContent) })

	// /body
	e.GET("/body", getBody)
}
