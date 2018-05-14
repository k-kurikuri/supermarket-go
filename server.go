package main

import (
	"html/template"
	"net/http"

	"github.com/k-kurikuri/supermarket-go/app"
	"github.com/labstack/echo"
)

const (
	port = ":8000"
)

var (
	e *echo.Echo
)

type Result struct {
	Success bool `json:"success"`
}

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

	e.POST("/add", func(context echo.Context) error {
		// TODO: mongoにRequestパラメータ入れる
		return context.JSON(http.StatusOK, &Result{Success: true})
	})

	e.Logger.Fatal(e.Start(port))
}

// parse rendering files
func setRenderer() {
	renderer := &app.TemplateRender{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
}
