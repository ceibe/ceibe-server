package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type File struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Hash string        `bson:"hash"`
	Path string        `bson:"path"`
}

type ImageFile struct {
	File

	Name       string `bson:"name"`
	Resolution Size2D `bson:"resolution"`
	Format     string `bson:"format"`
	Rotation   uint8
}

type VideoFile struct {
	File

	Name       string        `bson:"name"`
	Resolution Size2D        `bson:"resolution"`
	Duration   time.Duration `bson:"duration"`
	Format     string        `bson:"format"`
	Chapters   []Chapter     `bson:"chapters"`
	Tracks     []Track       `bson:"tracks"`
}

type TrackType uint8

const (
	VideoTrackType    TrackType = 0
	AudioTrackType    TrackType = 1
	SubtitleTrackType TrackType = 2
)

type Track struct {
	Index    uint      `bson:"index"`
	Type     TrackType `bson:"type"`
	Language string    `bson:"language"`
	Codec    string    `bson:"codec"`
}

type Chapter struct {
	Name      string        `bson:"name"`
	Timestamp time.Duration `bson:"time"`
}
