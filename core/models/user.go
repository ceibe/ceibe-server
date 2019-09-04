package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type User struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}
