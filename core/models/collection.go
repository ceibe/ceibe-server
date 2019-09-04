package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type CollectionType uint8

const (
	GenericCollectionType CollectionType = 0
	PersonCollectionType  CollectionType = 1
	AlbumCollectionType   CollectionType = 2
	BookCollectionType    CollectionType = 3
)

type CollectionResource struct {
	Resource

	Index uint `bson:"index"`
}

type Collection struct {
	Id          bson.ObjectId    `bson:"_id,omitempty"`
	Type        CollectionType   `bson:"type"`
	Name        string           `bson:"name"`
	Description string           `bson:"description"`
	PublishDate time.Time        `bson:"publishDate"`
	Tags        []string         `bson:"tags"`
	CreatedAt   time.Time        `bson:"createdAt"`
	Images      CollectionImages `bson:"images"`
}

type CollectionImages struct {
	Poster     bson.ObjectId `bson:"poster"`
	Background bson.ObjectId `bson:"background"`
}

// Main collections for a person. All person has one.
type PersonCollection struct {
	Collection

	Person Person `bson:"person"`
}

type CollectionShares struct {
	User   bson.ObjectId `bson:"user"`
	Modify bool          `bson:"modify"`
	Remove bool          `bson:"remove"`
}
