package controller

import (
	"context"

	"../database"
	"../model"
	"github.com/mongodb/mongo-go-driver/mongo"
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
// func GetTestContents(col *string) (*model.JSON, error) {

// 	var filter interface{}
// 	s, err := db.Find(col, filter)

// 	return s, err
// }

// Find : Controller for retrieving data in the test collection
func Find(col string) (interface{}, error) {

	var filter interface{}
	var cur mongo.Cursor

	cur, err := db.Find(cur, "Books", filter)

	defer cur.Close(context.Background())

	// var result model.JSON
	var result []model.Book
	for cur.Next(context.Background()) {
		var raw model.Book
		if err := raw.FromBSON(cur); err != nil {
			return result, err
		}

		result = append(result, raw)
	}

	return result, err
}
