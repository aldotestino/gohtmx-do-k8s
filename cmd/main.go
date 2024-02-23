package main

import (
	"gohtmx/handler"
	"gohtmx/model"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	localChatStore := model.NewLocalChatStore()

	userHandler := handler.UserHandler{}
	chatHandler := handler.ChatHandler{
		Store: localChatStore,
	}

	app.GET("/", userHandler.HandleUserLoginPage)
	app.POST("/login", userHandler.HandleUserLogin)

	app.GET("/chat", chatHandler.HandleChatPage)
	app.POST("/chat", chatHandler.HandleCreateMessage)

	app.Start(":80")
}
