package db

import (
	// "fmt"
	"github.com/globalsign/mgo"
	"log"
)

type DB struct{}

const (
	//DbUser ...
	// DbUser = "username"
	//DbPassword ...
	// DbPassword = "password"
	//DbName ...
	DbName = "ibms"
)

var db *mgo.Session

//Init ...
func Init() {

	var err error
	db, err = ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectMongoDB ...
func ConnectDB() (*mgo.Session, error) {
	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		return nil, err
	}
	// defer session.Close()

	return session, nil
}

//GetSession ...
func GetSession() *mgo.Session {
	return db
}

//GetDB ...
func GetDB() *mgo.Database {
	return db.DB(DbName)
}
