package main

import (
	Phasmo "GameSite/GameSpec/Phasmo"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "public")
	e.Renderer = newTemplate()

	e.GET("/", gameSelect)
	e.GET("/Phasmo", phasmoPage)
	e.GET("/Terraria", terrPage)

	e.POST("/PhasmoStart", phasmoSession)

	e.Logger.Fatal(e.Start(":4242"))
}

func gameSelect(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func phasmoPage(c echo.Context) error {

	PhasmoBaseInfo := Phasmo.GetBaseInfo()
	return c.Render(http.StatusOK, "Phasmo", PhasmoBaseInfo)
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
		templates: template.Must(template.ParseGlob("public/view/html/*.html")),
	}
}

func phasmoSession(c echo.Context) error {
	return c.Render(http.StatusOK, "PhasmoGameSession", nil)
}
