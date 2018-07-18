package app

import (
	"os"

	"gopkg.in/mgo.v2"
)

// DBName Mongo Database Name
const DBName = "supermarket-go"

// Table Mongo Table Name
const Table = "todolist"

// GetSession get mongodb session
func GetSession() *mgo.Session {
	session, err := mgo.Dial(GetDial())
	if err != nil {
		panic(err)
	}

	return session
}

func GetDial() string {
	dial := os.Getenv("MONGO_DIAL")
	if dial == "" {
		dial = "127.0.0.1:27017"
	}

	return dial
}
