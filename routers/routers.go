package routers

import (
	"github.com/gorilla/mux"
	// "crudtestgo/controllers"
)

var Routes = func(router *mux.Router) {

	router.HandleFunc("/users", controllers.getUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.getUser).Methods("GET")
	router.HandleFunc("/user/create", controllers.createCuser).Methods("POST")
	router.HandleFunc("/user/update", controllers.updateUser).Methods("PUT")
	router.HandleFunc("/user/delete", controllers.deleteUser).Methods("DELETE")
}
