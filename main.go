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
		Templates: template.Must(parseAssets("views/index.html")),
	}
	e.Renderer = renderer
}

func parseAssets(filenames ...string) (*template.Template, error) {
	var ns = template.New("complex")

	for _, filename := range filenames {

		src, err := Asset(filename)
		if err != nil {
			return nil, err
		}

		s := string(src)
		name := filepath.Base(filename)

		_, err = ns.New(name).Parse(s)
		if err != nil {
			return nil, err
		}
	}

	return ns, nil
}
