package main

/*
Listening to port :4242


*/

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
	Ghosts := Phasmo.GetGhosts()

	//Phasmo HTTP requests
	e.GET("/Phasmo", func(c echo.Context) error {
		PhasmoBaseInfo := Phasmo.GetBaseInfo()
		return c.Render(http.StatusOK, "Phasmo", PhasmoBaseInfo)
	})

	e.POST("/PhasmoGhostSearch", func(c echo.Context) error {

		return c.Render(http.StatusOK, "PhasmoGameSession", nil)
	})

	//Terraria HTTP requests
	e.GET("/Terraria", terrPage)

	e.Logger.Fatal(e.Start(":4242"))
}

// Home page server components
func gameSelect(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

// Terraria server components
func terrPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

// Core server components
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
