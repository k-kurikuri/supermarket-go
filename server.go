package main

import (
	"html/template"
	"net/http"

	"path/filepath"

	"github.com/k-kurikuri/supermarket-go/app"
	"github.com/k-kurikuri/supermarket-go/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2/bson"
)

const (
	port = ":8000"
)

var (
	e *echo.Echo
)

// Result Response
type Result struct {
	Success bool `json:"success"`
}

// init initialize
func init() {
	e = echo.New()

	e.Static("public/css", filepath.Join("public", "css"))
	e.Static("public/tag", filepath.Join("public", "tag"))
	setRenderer()
}

// main main function
func main() {
	setRouter()

	e.Logger.Fatal(e.Start(port))
}

// setRouter URL routing
func setRouter() {
	e.GET("/", func(context echo.Context) error {
		cookie, err := app.GetUidCookie(context)
		if err != nil {
			cookie = app.CreateUidCookie(context)
		}

		// Mongodb Session
		session := app.GetSession()
		defer session.Close()

		collect := session.DB(app.DBName).C(app.Table)

		todoList := &model.TodoList{}
		err = collect.Find(bson.M{"uid": cookie.Value}).One(&todoList)
		if err != nil {
			log.Error(err)
		}

		return context.Render(http.StatusOK, "index.html", map[string]interface{}{
			"TodoList": todoList,
		})
	})

	e.POST("/add", func(context echo.Context) error {
		// TODO: mongoにRequestパラメータ入れる
		return context.JSON(http.StatusOK, &Result{Success: true})
	})
}

// setRenderer parse rendering files
func setRenderer() {
	renderer := &app.TemplateRender{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
}
