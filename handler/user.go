package handler

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h *UserHandler) HandleUserLoginPage(c echo.Context) error {
	cookie, err := c.Cookie("user")

	if err != nil || cookie.Value == "" {
		tmpl := template.Must(template.ParseFiles("view/index.html"))
		return tmpl.Execute(c.Response(), nil)
	}

	return c.Redirect(http.StatusFound, "/chat")
}

func (h *UserHandler) HandleUserLogin(c echo.Context) error {
	user := c.FormValue("user")

	user = strings.Trim(user, " ")
	user = template.HTMLEscapeString(user)

	if user == "" {
		return c.Redirect(http.StatusFound, "/")
	}

	c.SetCookie(&http.Cookie{
		Name:     "user",
		Value:    user,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.Redirect(http.StatusFound, "/chat")
}
