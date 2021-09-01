package main

import (
	"github.com/aaaaaaaaaaa/config"
	"github.com/aaaaaaaaaaa/route"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){

	log.Println("Server started on: http://localhost:8000")
	config.CreateDatabase()
	route1 := mux.NewRouter()
	route.LoadingRoutes(route1)
	log.Fatal(http.ListenAndServe(":8000", route1))
}
