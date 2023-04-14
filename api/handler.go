package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getBody(c echo.Context) error  {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, json_map)
}
