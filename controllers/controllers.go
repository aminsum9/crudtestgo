package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"isbn"`
	Email    string `json:"title"`
	Password string `json:"password"`
}

var users []User

func TesConnection(ress http.ResponseWriter, router *http.Request) {

	type Tes struct {
		Success bool
	}

	var tes Tes

	tes.Success = true

	json.NewEncoder(ress).Encode(tes)
}

func GetUsers(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ress).Encode(users)
}

func GetUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)

	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(ress).Encode(item)
			return
		}
	}
	json.NewEncoder(ress).Encode(&User{})
}

func CreateUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	var user User
	router.ParseForm()

	_ = json.NewDecoder(router.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(10000000))
	users = append(users, user)
	json.NewEncoder(ress).Encode(user)
}

func UpdateUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(router.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(ress).Encode(user)
			return
		}
	}
}

func DeleteUser(ress http.ResponseWriter, router *http.Request) {

	ress.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(ress).Encode(users)
}
