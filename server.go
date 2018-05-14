package main

import (
	"html/template"
	"net/http"

	"github.com/k-kurikuri/supermarket-go/app"
	"github.com/labstack/echo"
)

var e *echo.Echo

func init() {
	e = echo.New()

	e.Static("public/css", "public/css")
	e.Static("public/tag", "public/tag")
	setRenderer()
}

func main() {
	e.GET("/", func(context echo.Context) error {
		if _, err := app.GetUidCookie(context); err != nil {
			app.SetUidCookie(context)
		} else {
			//todo: cookie.Value
		}

		return context.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.Logger.Fatal(e.Start(":8000"))
}

// parse rendering files
func setRenderer() {
	renderer := &app.TemplateRender{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
}
