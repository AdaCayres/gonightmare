package route


import (
	"github.com/aaaaaaaaaaa/controller"
	"github.com/gorilla/mux"
)

func LoadingRoutes(route *mux.Router){
	route.HandleFunc("/stock", controller.GetAll).Methods("GET")
	route.HandleFunc("/stock/{id}", controller.GetOne).Methods("GET")
	route.HandleFunc("/stock", controller.Insert).Methods("POST")
	route.HandleFunc("/stock/{id}", controller.Delete).Methods("DELETE")
	route.HandleFunc("/stock/{id}", controller.Update).Methods("PUT")
}