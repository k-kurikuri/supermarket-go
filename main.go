package main

import (
	"html/template"
	"path/filepath"

	"github.com/k-kurikuri/supermarket-go/app"
	"github.com/labstack/echo"
)

const (
	port = ":8000"
)

var (
	e *echo.Echo
)

// init initialize
func init() {
	e = echo.New()

	e.Static("public/css", filepath.Join("public", "css"))
	e.Static("public/tag", filepath.Join("public", "tag"))
	setRenderer()
}

// main main function
func main() {
	router := app.NewRouter()
	router.Boot(e)

	e.Logger.Fatal(e.Start(port))
}

// setRenderer parse rendering files
func setRenderer() {
	renderer := &app.TemplateRender{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
}
