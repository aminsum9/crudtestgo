package main

import (
	"log"
	"net/http"

	"github.com/aminsum9/test-crud/mux"
	"github.com/aminsum9/test-crud/routes"
)

func main() {
	//implementing router
	router := mux.NewRouter()

	//sending router to a different package named "routes"
	routes.BasicCrudRoutes(router)

	//this server runs here
	log.Fatal(http.ListenAndServe(":8000", router))
}
