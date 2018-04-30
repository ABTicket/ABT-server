package main

import (
	"fmt"
	"net/http"

	"./controllers"
)

func main() {
	router := controllers.NewRouter()
	http.ListenAndServe(":8080", router)
	fmt.Println("start server...")
}
