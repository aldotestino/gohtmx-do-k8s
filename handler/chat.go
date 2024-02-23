package handler

import (
	"gohtmx/model"
	"net/http"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	Store model.ChatStore
}

func (h *ChatHandler) HandleChatPage(c echo.Context) error {
	cookie, err := c.Cookie("user")

	if err != nil || cookie.Value == "" {
		return c.Redirect(http.StatusFound, "/")
	}

	messages := h.Store.GetAll()

	tmpl := template.Must(template.ParseFiles("view/chat.html"))
	return tmpl.Execute(c.Response(), map[string][]*model.Message{
		"Messages": messages,
	})
}

func (h *ChatHandler) HandleCreateMessage(c echo.Context) error {

	cookie, err := c.Cookie("user")
	if err != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	content := c.FormValue("content")
	content = strings.Trim(content, " ")

	message := model.NewMessage(cookie.Value, content)
	createdMessage, _ := h.Store.Create(message)

	tmpl := template.Must(template.ParseFiles("view/chat.html"))
	return tmpl.ExecuteTemplate(c.Response(), "message-item", createdMessage)
}
