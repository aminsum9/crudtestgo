package routers

import (
	"crudtestgo/controllers"

	"github.com/gorilla/mux"
)

var Routes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.TesConnection).Methods("GET")

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/update", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/delete", controllers.DeleteUser).Methods("DELETE")
}
