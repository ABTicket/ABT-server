package models

import (
	"gopkg.in/mgo.v2/bson"
	// "time"
)

type FilmShow struct {
	Id         bson.ObjectId `bson:"_id"`
	FilmName   string        `bson:"filmName"`
	CinemaName string        `bson:"cinemaName"`
	Time       string        `bson:"time"`
	// Time time.Time     `bson:"time"`
}
