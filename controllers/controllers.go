package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"database/sql"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
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

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-test")

	resultData, err := db.Query("SELECT * FROM users")

	type RessponseData struct {
		Success bool
		User    []User
		Error   string
	}

	var result []User

	if err != nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, User: result, Error: err.Error()})
	}

	defer resultData.Close()

	for resultData.Next() {
		var id int
		var name string
		var email string
		var password string

		err = resultData.Scan(&id, &name, &email, &password)

		if err != nil {
			json.NewEncoder(ress).Encode(RessponseData{Success: false, User: result, Error: err.Error()})
		}

		result = append(result, User{ID: id, Name: name, Email: email})
	}

	if len(result) >= 1 {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, User: result})
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, User: result})
	}
}

func GetUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)

	type RessponseData struct {
		Success bool
		User    User
		Error   string
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-test")
	if err != nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, User: User{}, Error: ""})
	}

	idUser, err := strconv.Atoi(params["id"])

	resultData, err := db.Query("SELECT * FROM `crud-test`.users WHERE id = ?", idUser)

	defer resultData.Close()

	resultData.Next()

	var id int
	var name string
	var email string
	var password string

	resultData.Scan(&id, &name, &email, &password)

	responseData := User{ID: id, Name: name, Email: email, Password: password}

	if responseData.ID != 0 {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, User: responseData, Error: ""})
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, User: responseData, Error: "User tidak ditemukan!"})
	}

}

func CreateUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	var user User

	decoder := json.NewDecoder(router.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(ress, err.Error(), http.StatusBadRequest)
		return
	}

	name := user.Name
	email := user.Email
	password := user.Password

	type RessponseData struct {
		Success bool
		Message string
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-test")

	if err != nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, Message: err.Error()})
	}

	_, err = db.Query("INSERT INTO `crud-test`.users VALUES(0,?,?,?)", name, email, password)

	if err == nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, Message: "Berhasil menambah user!"})
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, Message: "Gagal menambah user!"})
	}
}

func UpdateUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	var user User

	decoder := json.NewDecoder(router.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(ress, err.Error(), http.StatusBadRequest)
		return
	}

	idUser := user.ID
	name := user.Name
	email := user.Email
	password := user.Password

	type RessponseData struct {
		Success bool
		Message string
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-test")

	if err != nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, Message: err.Error()})
	}

	_, err = db.Query("UPDATE `crud-test`.users SET name = ?, email = ?, password = ? WHERE id = ?", name, email, password, idUser)

	if err == nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, Message: "Berhasil update data user!"})
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, Message: "Gagal update data user!"})
	}
}

func DeleteUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	var user User

	decoder := json.NewDecoder(router.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(ress, err.Error(), http.StatusBadRequest)
		return
	}

	idUser := user.ID

	type RessponseData struct {
		Success bool
		Message string
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-test")

	if err != nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, Message: err.Error()})
	}

	_, err = db.Query("DELETE FROM `crud-test`.users WHERE id = ?", idUser)

	if err == nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, Message: "Berhasil delete data user!"})
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, Message: "Gagal delete data user!"})
	}
}
