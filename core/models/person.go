package models

import "github.com/globalsign/mgo/bson"

type PersonGender uint8

const (
	MaleGender   PersonGender = 0
	FemaleGender PersonGender = 1
)

type Person struct {
	Id bson.ObjectId `bson:"_id,omitempty"`

	Name           string       `bson:"name"`
	PatronymicName string       `bson:"patronymicName"`
	Summary        string       `bson:"summary"`
	Gender         PersonGender `bson:"gender"`
	Address        string       `bson:"address"`
	Tags           []string     `bson:"tags"`
}
