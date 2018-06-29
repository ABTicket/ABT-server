package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	. "../log"
	. "../models"
	"../utils"
)

func FilmShowGetOne(w http.ResponseWriter, r *http.Request) {
	// GET, 从URL中读取参数, 直接使用mux.Vars(r)
	vars := mux.Vars(r)
	filmShowId := vars["filmShowId"]

	filmShow := FilmShow{}
	err := Db["filmShows"].FindId(bson.ObjectIdHex(filmShowId)).One(&filmShow)
	if err != nil {
		Log.Errorf("Get filmShow id: %s failed, %v", filmShowId, err)
		utils.FailureResponse(&w, "获取放映信息失败", "")
		return
	}

	Log.Noticef("Get filmShow successfully: %s", filmShow)
	utils.SuccessResponse(&w, "获取放映信息成功", filmShow)
}

func FilmShowGetAll(w http.ResponseWriter, r *http.Request) {
	var filmShows []FilmShow
	err := Db["filmShows"].Find(nil).All(&filmShows)
	if err != nil {
		Log.Errorf("get all filmShows failed, %v", err)
		utils.FailureResponse(&w, "获取放映列表失败", "")
		return
	}
	Log.Notice("get all filmShow successfully")
	utils.SuccessResponse(&w, "获取放映列表成功", filmShows)
}

// func FilmShowGetFilms(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	cinemaId := vars["cinemaId"]

// 	var flimList []string
// 	err := Db["filmShows"].Find(bson.M{"cinemaName": cinemaName}).Distinct("filmName", &flimList)
// 	if err != nil {
// 		Log.Errorf("get flimList failed, %v", err)
// 		utils.FailureResponse(&w, "获取电影列表失败", "")
// 		return
// 	}
// 	Log.Notice("get flimList successfully")
// 	utils.SuccessResponse(&w, "获取电影列表成功", flimList)
// }

// func FilmShowGetCinemas(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	filmName := vars["filmName"]

// 	var cinemaList []string
// 	err := Db["filmShows"].Find(bson.M{"filmName": filmName}).Distinct("cinemaName", &cinemaList)
// 	if err != nil {
// 		Log.Errorf("get cinemaList failed, %v", err)
// 		utils.FailureResponse(&w, "获取电影院列表失败", "")
// 		return
// 	}
// 	Log.Notice("get cinemaList successfully")
// 	utils.SuccessResponse(&w, "获取电影院列表成功", cinemaList)
// }

func FilmShowAddOne(w http.ResponseWriter, r *http.Request) {
	// 1. load request's body
	vars := mux.Vars(r)
	filmId := vars["filmId"]
	cinemaId := vars["cinemaId"]

	newFilmShow := FilmShow{}
	ok := utils.LoadRequestBody(r, "insert filmShow", &newFilmShow)
	if !ok {
		utils.FailureResponse(&w, "新建放映失败", "")
		return
	}
	newFilmShow.FilmId = bson.ObjectIdHex(filmId)
	newFilmShow.CinemaId = bson.ObjectIdHex(cinemaId)

	// 2. verify the film existed or not
	existedFilm := Film{}
	err := Db["films"].FindId(newFilmShow.FilmId).One(&existedFilm)
	if err != nil {
		Log.Errorf("insert filmShow failed: film is not existed")
		utils.FailureResponse(&w, "电影不存在", "")
		return
	}
	// 3. verify the film existed or not
	existedCinema := Cinema{}
	err = Db["cinemas"].FindId(newFilmShow.CinemaId).One(&existedCinema)
	if err != nil {
		Log.Errorf("insert filmShow failed: cinema is not existed")
		utils.FailureResponse(&w, "电影院不存在", "")
		return
	}
	// 4. verify the filmShow existed or not
	existedFilmShow := FilmShow{}
	err = Db["filmShows"].
		Find(bson.M{"filmId": newFilmShow.FilmId, "cinemaId": newFilmShow.CinemaId, "time": newFilmShow.Time}).
		One(&existedFilmShow)
	if err == nil {
		Log.Errorf("insert filmShow failed: filmShow is existed")
		utils.FailureResponse(&w, "放映已存在", "")
		return
	}
	// 5. set a new id
	newFilmShow.Id = bson.NewObjectId()
	// 6. insert into db
	err = Db["filmShows"].Insert(&newFilmShow)
	if err != nil {
		Log.Error("insert filmShow falied: insert into db failed, ", err)
		utils.FailureResponse(&w, "添加放映失败", "")
		return
	}
	// 7. success
	Log.Notice("add one filmShow successfully")
	utils.SuccessResponse(&w, "添加放映成功", "")
}

func FilmShowUpdateOne(w http.ResponseWriter, r *http.Request) {
	// 1. 获得URL中的参数
	vars := mux.Vars(r)
	filmShowId := vars["filmShowId"]
	filmId := vars["filmId"]
	cinemaId := vars["cinemaId"]
	// 2. 从request中解析出body数据
	newFilmShow := FilmShow{}
	ok := utils.LoadRequestBody(r, "update filmShow", &newFilmShow)
	if !ok {
		utils.FailureResponse(&w, "修改放映信息失败", "")
	}
	newFilmShow.Id = bson.ObjectIdHex(filmShowId)
	newFilmShow.FilmId = bson.ObjectIdHex(filmId)
	newFilmShow.CinemaId = bson.ObjectIdHex(cinemaId)

	// 3. 修改数据
	// convert structure to bson.M, used to update
	updateData, _ := bson.Marshal(&newFilmShow)
	updateFilmShow := bson.M{}
	_ = bson.Unmarshal(updateData, &updateFilmShow)

	// 此处更新时如果没有"$set",会将整行直接覆盖，而不是按需修改
	err := Db["filmShows"].Update(bson.M{"_id": newFilmShow.Id}, bson.M{"$set": updateFilmShow})
	if err != nil {
		Log.Error("update filmShow failed: failed to update data into db, ", err)
		utils.FailureResponse(&w, "修改放映信息失败", "")
		return
	}
	// 4. 成功返回
	Log.Notice("update filmShow successfully")
	utils.SuccessResponse(&w, "修改放映成功", "")
}

func FilmShowDeleteOne(w http.ResponseWriter, r *http.Request) {
	filmShowId := mux.Vars(r)["filmShowId"]

	err := Db["filmShows"].Remove(bson.M{"_id": bson.ObjectIdHex(filmShowId)})
	if err != nil {
		Log.Error("delete filmShow from db failed: ", err)
		utils.FailureResponse(&w, "删除放映失败", "")
		return
	}

	Log.Notice("delete filmShow successfully")
	utils.SuccessResponse(&w, "删除放映成功", "")
}

var FilmShowRoutes Routes = Routes{
	Route{"FilmShowGetOne", "GET", "/filmShow/{filmShowId}", FilmShowGetOne},
	Route{"FilmShowGetAll", "GET", "/filmShow", FilmShowGetAll},
	//Route{"FilmShowGetFilms", "GET", "/filmShow/cinema/{cinemaName}", FilmShowGetFilms},
	//Route{"FilmShowGetCinemas", "GET", "/filmShow/film/{filmName}", FilmShowGetCinemas},
	Route{"FilmShowAddOne", "POST", "/filmShow/", FilmShowAddOne},
	Route{"FilmShowUpdateOne", "PUT", "/filmShow/{filmShowId}", FilmShowUpdateOne},
	Route{"FilmShowDeleteOne", "DELETE", "/filmShow/{filmShowId}", FilmShowDeleteOne},
}
