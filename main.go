package testcrud

import (
	"crudtestgo/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//implementing router
	router := mux.NewRouter()

	//sending router to a different package named "routes"
	routers.Routes(router)

	//this server runs here
	log.Fatal(http.ListenAndServe(":8000", router))
}
