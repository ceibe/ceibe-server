package models

type GeoPosition struct {
	Latitude  float32 `bson:"lat"`
	Longitude float32 `bson:"lng"`
}
