package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Resource struct {
	Id bson.ObjectId `bson:"_id,omitempty"`

	Name           string        `bson:"name"`
	Visualizations uint64        `bson:"visualizations"`
	AuthorId       bson.ObjectId `bson:"user"`
	CreatedAt      time.Time     `bson:"createdAt"`
	Where          GeoPosition   `bson:"where"`
	Tags           []string      `bson:"tags"`
}

type ImageResource struct {
	Resource
}

type VideoResource struct {
	Resource
}
