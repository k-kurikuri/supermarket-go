package app

import (
	"log"
	"net/http"

	"github.com/k-kurikuri/supermarket-go/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) Index(c echo.Context) error {
	cookie, err := GetUidCookie(c)
	if err != nil {
		cookie = CreateUidCookie(c)
	}

	// Mongodb Session
	session := GetSession()
	defer session.Close()

	collect := session.DB(DBName).C(Table)

	todoList := &model.TodoList{}
	collect.Find(bson.M{"uid": cookie.Value}).One(&todoList)

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"TodoList": todoList,
	})
}

func (controller *Controller) AddTask(c echo.Context) error {
	todo := &model.Todo{Done: false}
	if err := c.Bind(todo); err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, &model.Result{Success: false})
	}

	cookie, err := GetUidCookie(c)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, &model.Result{Success: false, Message: err.Error()})
	}

	session := GetSession()
	defer session.Close()

	collect := session.DB(DBName).C(Table)
	selector := bson.M{"uid": cookie.Value}
	update := bson.M{"$push": bson.M{"todo": todo}}
	_, err = collect.Upsert(selector, update)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, &model.Result{Success: false, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &model.Result{Success: true})
}

func (controller *Controller) DeleteTask(c echo.Context) error {
	indexes := &model.Indexes{}
	if err := c.Bind(indexes); err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, &model.Result{Success: false})
	}

	cookie, err := GetUidCookie(c)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, &model.Result{Success: false, Message: err.Error()})
	}

	session := GetSession()
	defer session.Close()

	collect := session.DB(DBName).C(Table)
	selector := bson.M{"uid": cookie.Value}

	// 一度、配列内のデータをnilに設定し、その後nilをまとめてpullするハック Not Atomic..
	for _, delIndex := range indexes.Indexes {
		update := bson.M{"$set": bson.M{"todo." + delIndex: nil}}
		err = collect.Update(selector, update)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusOK, &model.Result{Success: false, Message: err.Error()})
		}
	}

	update := bson.M{"$pull": bson.M{"todo": nil, "multi": true}}
	err = collect.Update(selector, update)

	return c.JSON(http.StatusOK, map[string]interface{}{"delete": "OK"})
}

func (controller *Controller) PutTask(c echo.Context) error {
	putReq := &model.PutRequest{}
	if err := c.Bind(putReq); err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, &model.Result{Success: false})
	}

	cookie, err := GetUidCookie(c)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, &model.Result{Success: false, Message: err.Error()})
	}

	session := GetSession()
	defer session.Close()

	collect := session.DB(DBName).C(Table)
	selector := bson.M{"uid": cookie.Value}
	update := bson.M{"$set": bson.M{"todo." + putReq.Index: bson.M{"title": putReq.Todo.Title, "done": putReq.Todo.Done}}}
	err = collect.Update(selector, update)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusOK, &model.Result{Success: false, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"update": "OK"})
}
