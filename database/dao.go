package database

import (
	"context"
	"log"
	"strconv"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// DAO : Data Access Object
type DAO struct {
	Server   string
	Database string
	User     string
	Password string
}

var db *mongo.Database

// Connect : Establishes connection with MongoDB
func (d *DAO) Connect() {

	uri := "mongodb://" + d.User + ":" + d.Password + "@" + d.Server
	client, err := mongo.Connect(context.Background(), uri, nil)

	if err != nil {
		panic(err)
	}

	db = client.Database(d.Database)

	log.Println("Establishing connection with mongoDB - " + "'" + d.Database + "'")
}

// CollectionSize : Returns the size of a collection
func (d *DAO) CollectionSize(n *string) (string, error) {

	var i interface{}
	col := db.Collection(*n)
	num, err := col.Count(nil, i)

	if err != nil {
		return "0", err
	}

	s := strconv.FormatInt(num, 10)

	return s, err
}

// // Find : Runs a find function on the collection w/ filter
// func (d *DAO) Find(collectionName *string, filter interface{}) (*model.JSON, error) {

// 	var cur mongo.Cursor
// 	ctx := context.Background()

// 	col := db.Collection(*collectionName)
// 	cur, err := col.Find(ctx, filter)

// 	defer cur.Close(ctx)

// 	var result model.JSON
// 	doc := bson.NewDocument()
// 	for cur.Next(ctx) {
// 		doc.Reset()
// 		if err := cur.Decode(doc); err != nil {
// 			log.Fatal(err)
// 		}

// 		var raw model.Document
// 		if err := json.Unmarshal([]byte(doc.ToExtJSON(false)), &raw); err != nil {
// 			log.Fatal(err)
// 		}
// 		result = append(result, raw)

// 	}

// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return &result, err
// }

// Find : Runs a find function on the collection w/ filter
func (d *DAO) Find(cur mongo.Cursor, collectionName string, filter interface{}) (mongo.Cursor, error) {
	col := db.Collection(collectionName)
	cur, err := col.Find(nil, filter)

	return cur, err
}

// // EstimatedDocumentCount : Gets an estimate of the count of documents
// func (d *DAO) EstimatedDocumentCount(collectionName *string) (*int64, error) {

// 	result := db.Collection(*collectionName)
// 	count, err := result.EstimatedDocumentCount(nil)

// 	return count, err
// }
