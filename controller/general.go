package controller

import (
	"../database"
	"../model"
)

var db = database.DAO{}

func init() {
	db.Server = "127.0.0.1:27017"
	db.Database = "chris_db"
	db.User = "user"
	db.Password = "password"
	db.Connect()
}

// GetCollectionSize : Controller for retrieving the size of a collection
func GetCollectionSize(col *string) (string, error) {
	s, err := db.CollectionSize(col)

	return s, err
}

// GetTestContents : Controller for retrieving data in the test collection
func GetTestContents(col *string) (*model.JSON, error) {

	var filter interface{}
	s, err := db.Find(col, filter)

	return s, err
}
