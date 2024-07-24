package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Static("/public"))
	e.Renderer = newTemplate()

	e.GET("/", gameSelect)
	e.PUT("/Phasmo", phasmoPage)
	e.PUT("/Terraria", terrPage)

	e.Logger.Fatal(e.Start(":4242"))
}

func gameSelect(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func phasmoPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func terrPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("public/view/*.html")),
	}
}
