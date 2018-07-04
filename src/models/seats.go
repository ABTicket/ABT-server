package models

import (
	"gopkg.in/mgo.v2/bson"
	// "time"
)

type Seats struct {
	Id         bson.ObjectId `bson:"_id"`
	FilmShowId bson.ObjectId `bson:"filmShowId"`
	SeatsSold  []int         `bson:"seatsSold"`
	// 座位号默认1——80，数据库只保存已售出的座位号，未售出的不保存
}
