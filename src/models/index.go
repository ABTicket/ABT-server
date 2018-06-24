package models

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"

	"../configs"
	. "../log"
)

var Db = make(map[string]*mgo.Collection)
var modelsList = []string{
	"users", "films", "cinemas", "filmShows", "Imgs",
}

func DbInit() {
	session, err := mgo.Dial(configs.MONGODB_URL)
	Log.Notice("connecting to database...")
	if err != nil {
		Log.Error("connecting to database failed...")
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	for _, model := range modelsList {
		Db[model] = session.DB(configs.MONGODB_NAME).C(model)
	}
}
