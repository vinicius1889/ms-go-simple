package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"./controller"
)


func main(){
	route := mux.NewRouter()
	route.HandleFunc("/call", controller.SetLastCall)
	route.HandleFunc("/calls/{key}/{limit}", controller.GetLastCalls)
	route.HandleFunc("/health", controller.HealthFunc)
	http.ListenAndServe(":8089",route)
}