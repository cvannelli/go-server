package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Book : Model for a book
type Book struct {
	ID    objectid.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title string            `json:"title"  bson:"title"`
	Year  int               `json:"year"  bson:"year"`
}

type Resource interface{}

func (s *Book) FromBSON(cur mongo.Cursor) error {
	doc := bson.NewDocument()
	if err := cur.Decode(doc); err != nil {
		return err
	}
	u, err := doc.MarshalBSON()
	if err != nil {
		return err
	}
	err = bson.Unmarshal(u, s)

	return err
}
