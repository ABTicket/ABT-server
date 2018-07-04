package main

import (
	"fmt"
	"net/http"

	"./configs"
	"./controllers"
	"./models"
)

func init() {
	models.DbInit()
}

func main() {
	router := controllers.NewRouter()
	addr := configs.HOST + ":" + configs.PORT
	fmt.Printf("start server at %v...\n", addr)
	http.ListenAndServe(addr, router)
}
