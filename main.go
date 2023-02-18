package main

import (
	db "crudtestgo/db"
	"crudtestgo/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Db()

	router := mux.NewRouter()

	routers.Routes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
