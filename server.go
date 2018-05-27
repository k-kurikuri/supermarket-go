package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
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
	Success bool   `json:"success"`
	Message string `json:"message"`
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
	e.GET("/", func(c echo.Context) error {
		cookie, err := app.GetUidCookie(c)
		if err != nil {
			cookie = app.CreateUidCookie(c)
		}

		// Mongodb Session
		session := app.GetSession()
		defer session.Close()

		collect := session.DB(app.DBName).C(app.Table)

		todoList := &model.TodoList{}
		collect.Find(bson.M{"uid": cookie.Value}).One(&todoList)

		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"TodoList": todoList,
		})
	})

	e.POST("/add", func(c echo.Context) error {
		todo := &model.Todo{Done: false}
		if err := c.Bind(todo); err != nil {
			log.Print(err)
			return c.JSON(http.StatusOK, &Result{Success: false})
		}

		cookie, err := app.GetUidCookie(c)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusOK, &Result{Success: false, Message: err.Error()})
		}

		session := app.GetSession()
		defer session.Close()

		collect := session.DB(app.DBName).C(app.Table)
		selector := bson.M{"uid": cookie.Value}
		update := bson.M{"$push": bson.M{"todo": todo}}
		_, err = collect.Upsert(selector, update)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusOK, &Result{Success: false, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, &Result{Success: true})
	})

	e.PUT("/update", func(c echo.Context) error {
		s, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		var body map[string]interface{}
		if err := json.Unmarshal(s, &body); err != nil {
			return err
		}

		idx, _ := body["index"].(string)
		title, _ := body["title"].(string)
		done, _ := body["done"].(bool)
		todo := &model.Todo{Title: title, Done: done}

		cookie, err := app.GetUidCookie(c)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusOK, &Result{Success: false, Message: err.Error()})
		}

		session := app.GetSession()
		defer session.Close()

		collect := session.DB(app.DBName).C(app.Table)
		selector := bson.M{"uid": cookie.Value}
		update := bson.M{"$set": bson.M{"todo." + idx: bson.M{"title": todo.Title, "done": todo.Done}}}
		err = collect.Update(selector, update)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusOK, &Result{Success: false, Message: err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"update": "OK"})
	})
}

// setRenderer parse rendering files
func setRenderer() {
	renderer := &app.TemplateRender{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
}
