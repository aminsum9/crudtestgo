package routers

import (
	"crudtestgo/controllers/usercontroller"
	"crudtestgo/middleware/authmiddleware"

	"github.com/gorilla/mux"
)

var Routes = func(router *mux.Router) {

	router.Use(authmiddleware.AuthMiddleware)

	router.HandleFunc("/users", usercontroller.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", usercontroller.GetUser).Methods("GET")
	router.HandleFunc("/user/create", usercontroller.CreateUser).Methods("POST")
	router.HandleFunc("/user/update", usercontroller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/delete", usercontroller.DeleteUser).Methods("DELETE")
}
