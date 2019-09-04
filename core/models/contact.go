package models

import "github.com/globalsign/mgo/bson"

type ContactType uint8

const (
	EmailContactType    ContactType = 0
	HomepageContactType ContactType = 1
)

type Contact struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Person bson.ObjectId `bson:"person"`
	Type   ContactType   `bson:"type"`
}

type UrlContact struct {
	Contact

	Url string `bson:"url"`
}
