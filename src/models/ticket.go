package models

import (
	"gopkg.in/mgo.v2/bson"
	// "time"
)

type Ticket struct {
	Id         bson.ObjectId `bson:"_id"`
	FilmShowId bson.ObjectId `bson:"filmId"`
	Seat       int           `bson:"seat"`
}
