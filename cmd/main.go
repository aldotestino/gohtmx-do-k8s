package main

import (
	"gohtmx/handler"
	"gohtmx/model"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	localUserStore := model.NewLocalUserStore()

	userHandler := handler.UserHandler{
		Store: localUserStore,
	}

	app.GET("/", userHandler.HandlePage)
	app.POST("/user", userHandler.HandleAdd)
	app.Start(":80")
}
