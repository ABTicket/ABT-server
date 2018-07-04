package models

import (
	"gopkg.in/mgo.v2/bson"
	// "time"
)

type FilmShow struct {
	Id       bson.ObjectId `bson:"_id"`
	FilmId   string        `bson:"filmId"`
	CinemaId bson.ObjectId `bson:"cinemaId"`
	Time     int64         `bson:"time"` // 使用字符串容易设置成固定时间
}
