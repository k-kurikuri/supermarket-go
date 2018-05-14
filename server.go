package main

import (
	"html/template"
	"net/http"

	"io"
	"time"

	"github.com/labstack/echo"
	"github.com/rs/xid"
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

const (
	userCookieName = "_uid"
)

func init() {
	e = echo.New()

	e.Static("public/css", "public/css")
	e.Static("public/tag", "public/tag")
	setRenderer()
}

func main() {
	e.GET("/", func(context echo.Context) error {
		if _, err := getUidCookie(context); err != nil {
			setUidCookie(context)
		} else {
			//todo: cookie.Value
		}

		return context.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	e.Logger.Fatal(e.Start(":8000"))
}

// parse rendering files
func setRenderer() {
	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
}

// Cookieにuidを書き込む
func setUidCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = userCookieName
	cookie.Value = createUuid()
	cookie.Expires = time.Now().Add(2 * time.Hour)
	c.SetCookie(cookie)
}

// uidキーのCookieを取得する
func getUidCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie(userCookieName)

	return cookie, err
}

// uuidを生成する
func createUuid() string {
	uid := xid.New()

	return uid.String()
}
