package api

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func StartServer() {
	// new echo instance
	e := echo.New()
	// router
	router(e)
	// get server address
	address := ":" + viper.Get("server.port").(string)
	//log server address
	e.Logger.Fatal(e.Start(address))
}