package main

import (
	"html/template"
	"net/http"

	"io"

	"github.com/labstack/echo"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

var e *echo.Echo

func init() {
	e = echo.New()

	e.Static("public/css", "public/css")
	e.Static("public/tag", "public/tag")
	setRenderer()
}

func main() {
	e.GET("/", func(context echo.Context) error {
		return context.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.Logger.Fatal(e.Start(":8000"))
}

func setRenderer() {
	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer
}
