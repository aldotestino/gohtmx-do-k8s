package handler

import (
	"gohtmx/model"
	"text/template"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Store model.UserStore
}

func (h *UserHandler) HandlePage(c echo.Context) error {
	users := h.Store.FindAll()
	tmpl := template.Must(template.ParseFiles("view/index.html"))
	return tmpl.Execute(c.Response(), map[string][]*model.User{
		"Users": users,
	})
}

func (h *UserHandler) HandleAdd(c echo.Context) error {
	email := c.FormValue("email")
	user := model.NewUser(email)
	createdUser, _ := h.Store.Create(user)

	tmpl := template.Must(template.ParseFiles("view/index.html"))
	return tmpl.ExecuteTemplate(c.Response(), "list-item", createdUser)
}
