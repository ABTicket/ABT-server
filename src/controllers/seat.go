package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	. "../log"
	. "../models"
	"../utils"
)

/*
func SeatAddOne(newOrder Order) bool {
	newSeat := Seat{}
	newSeat.FilmShowId = newOrder.FilmShowId
	newSeat.SeatNum = newOrder.SeatNum

	newSeat.Id = bson.NewObjectId()

	err := Db["seats"].Insert(&newSeat)
	if err != nil {
		Log.Error("insert order falied: insert into db failed, ", err)
		return false
	}

	Log.Notice("add one seat successfully")
	return true
}

func SeatDeleteOne(deleteOrder Order) bool {
	deleteSeat := Seat{}
	deleteSeat.FilmShowId = deleteOrder.FilmShowId
	deleteSeat.SeatNum = deleteOrder.SeatNum

	err := Db["seats"].Remove(bson.M{"filmshowid": deleteSeat.FilmShowId, "seatNum": deleteSeat.SeatNum})
	if err != nil {
		Log.Error("delete seat from db failed: ", err)
		return false
	}

	Log.Notice("delete seat successfully")
	return true
}
*/

type addSeat struct {
	filmShowId string
	seats      []int
}

func SeatAdd(w http.ResponseWriter, r *http.Request) {
	newAddSeat := addSeat{}
	ok := utils.LoadRequestBody(r, "add seats", &newAddSeat)
	if !ok {
		utils.FailureResponse(&w, "添加座位失败", "")
		return
	}
	query := bson.M{"filmShowId": newAddSeat.filmShowId}
	update := bson.M{"$push": bson.M{"seatsSold": newAddSeat.seats}}
	err := Db["seats"].Update(query, update)
	if err != nil {
		Log.Errorf("添加座位失败：%v", err)
		utils.FailureResponse(&w, "添加座位失败", "")
	}
	Log.Errorf("添加座位成功", err)
	utils.FailureResponse(&w, "添加座位成功", "")
}

func SeatGetFromFilmShowId(w http.ResponseWriter, r *http.Request) {
	var seats Seats

	vars := mux.Vars(r)
	filmShowId := vars["filmShowId"]

	err := Db["seats"].Find(bson.M{"filmShowId": filmShowId}).One(&seats)
	if err != nil {
		if err.Error() == "not found" {
			newSeats := Seats{bson.NewObjectId(), bson.ObjectIdHex(filmShowId), []int{}}
			Db["seats"].Insert(newSeats)
			Log.Notice("get seatNums successfully")
			utils.SuccessResponse(&w, "获取座位号列表成功", newSeats)
		} else {
			Log.Errorf("get seatNums failed, %v", err)
			utils.FailureResponse(&w, "获取座位号列表失败", "")
			return
		}
	}
	Log.Notice("get seatNums successfully")
	utils.SuccessResponse(&w, "获取座位号列表成功", seats)
}

func SeatGetAll(w http.ResponseWriter, r *http.Request) {
	var seats []Seats
	err := Db["seats"].Find(nil).All(&seats)
	if err != nil {
		Log.Errorf("get all seats failed, %v", err)
		utils.FailureResponse(&w, "获取座位列表失败", "")
		return
	}
	Log.Notice("get all user successfully")
	utils.SuccessResponse(&w, "获取座位列表成功", seats)
}

var SeatRoutes Routes = Routes{
	// 添加座位号(已售出的), 请求体：{filmShowId: "", seats: [1, 2, 3]}
	Route{"SeatAdd", "POST", "/seat", SeatAdd},
	// 根据filmShowId(电影场次)获取这个场次对应的座位列表
	Route{"SeatGetFromFilmShowId", "GET", "/seat/{filmShowId}", SeatGetFromFilmShowId},
	Route{"SeatGetAll", "GET", "/seat", SeatGetAll},
}
