package models

import (
	"gopkg.in/mgo.v2/bson"
	// "time"
)

type Img struct {
	Id     bson.ObjectId `bson:"_id"`
	ImgUrl string        `bson:"imgUrl"`
}
