package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	. "../log"
	. "../models"
	"../utils"
)

func OrderAddOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	filmShowId := vars["filmShowId"]

	newOrder := Order{}
	ok := utils.LoadRequestBody(r, "insert order", &newOrder)
	if !ok {
		utils.FailureResponse(&w, "新建订单失败", "")
		return
	}
	newOrder.UserId = bson.ObjectIdHex(userId)
	newOrder.FilmShowId = bson.ObjectIdHex(filmShowId)

	existedOrder := Order{}
	err := Db["orders"].Find(bson.M{"filmShowId": newOrder.FilmShowId, "seatNum": newOrder.SeatNum}).One(&existedOrder)
	if err == nil {
		Log.Errorf("insert order failed: seat %d has been sold", newOrder.SeatNum)
		utils.FailureResponse(&w, "该场次此座位已售出", "")
		return
	}

	newOrder.Id = bson.NewObjectId()

	err = Db["orders"].Insert(&newOrder)
	if err != nil {
		Log.Error("insert order falied: insert into db failed, ", err)
		utils.FailureResponse(&w, "添加订单失败", "")
		return
	}

	Log.Notice("add one order successfully")
	utils.SuccessResponse(&w, "添加订单成功", "")
}

//func OrderUpdateOne(w http.ResponseWriter, r *http.Request) {}

func OrderDeleteOne(w http.ResponseWriter, r *http.Request) {
	orderId := mux.Vars(r)["orderId"]

	err := Db["orders"].Remove(bson.M{"_id": bson.ObjectIdHex(orderId)})
	if err != nil {
		Log.Error("delete order from db failed: ", err)
		utils.FailureResponse(&w, "删除订单失败", "")
		return
	}

	Log.Notice("delete order successfully")
	utils.SuccessResponse(&w, "删除订单成功", "")
}

func OrderGetFromUserId(w http.ResponseWriter, r *http.Request) {
	var orders []Order

	vars := mux.Vars(r)
	userId := vars["userId"]

	err := Db["orders"].Find(bson.M{"userId": bson.ObjectIdHex(userId)}).All(&orders)
	if err != nil {
		Log.Errorf("get orders failed, %v", err)
		utils.FailureResponse(&w, "获取订单列表失败", "")
		return
	}
	Log.Notice("get orders successfully")
	utils.SuccessResponse(&w, "获取订单列表成功", orders)
}

var OrderRoutes Routes = Routes{
	Route{"OrderAddOne", "POST", "/order/", OrderAddOne},
	Route{"OrderDeleteOne", "DELETE", "/order/", OrderDeleteOne},
	Route{"OrderGetFromUserId", "GET", "/order/{orderId}", OrderGetFromUserId},
}
