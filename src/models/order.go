package models

import (
	"gopkg.in/mgo.v2/bson"
	// "time"
)

type Order struct {
	Id         bson.ObjectId   `bson:"_id"`
	UserId     bson.ObjectId   `bson:"userId"`
	Tickets    []bson.ObjectId `bson:"tickets"`
	CreateTime int64           `bson:"createTime"`
}
