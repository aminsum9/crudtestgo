package testcrud

import (
	"log"
	"net/http"

	routers "github.com/aminsum9/testcrud/routers/routers"
	"github.com/gorilla/mux"
)

func main() {
	//implementing router
	router := mux.NewRouter()

	//sending router to a different package named "routes"
	routers.BasicCrudRoutes(router)

	//this server runs here
	log.Fatal(http.ListenAndServe(":8000", router))
}
