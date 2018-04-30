package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	. "../log"
)

func UserGetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	fmt.Fprintf(w, "UserGetOne: %v", userId)

	Log.Debugf("debug: %v", userId)
	Log.Info("info")
	Log.Notice("notice")
    Log.Warning("warning")
    Log.Error("err")
    Log.Critical("critival")
}

var UserRoutes Routes = Routes {
	Route {
		"UserGetOne",
		"GET",
		"/user/{userId}",
		UserGetOne,
	},
}

