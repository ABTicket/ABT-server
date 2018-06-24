package main

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"./configs"
	"./controllers"
	"./models"
)

func init() {
	// 已知bug，session中要保存自定义的结构体
	// 目前只能通过gob.Resister解决
	// see: https://github.com/astaxie/beego/issues/1110
	gob.Register(models.User{})

	models.DbInit()
}

func main() {
	router := controllers.NewRouter()
	http.ListenAndServe(configs.HOST+":"+configs.PORT, router)
	fmt.Println("start server...")
}
