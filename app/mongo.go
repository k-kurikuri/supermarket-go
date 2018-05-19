package app

import "gopkg.in/mgo.v2"

const dial = "127.0.0.1:27017"

// DBName Mongo Database Name
const DBName = "supermarket-go"

// Table Mongo Table Name
const Table = "todolist"

// GetSession get mongodb session
func GetSession() *mgo.Session {
	session, err := mgo.Dial(dial)
	if err != nil {
		panic(err)
	}

	return session
}
