package usercontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	db "crudtestgo/db"
	"crudtestgo/model/user"

	"github.com/gorilla/mux"
)

func GetUsers(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")

	log.Println(router.Header["Key"], " --------router")

	resultData, err := db.Db().Query("SELECT * FROM users")

	type RessponseData struct {
		Success bool
		User    []*user.User
		Error   string
	}

	var result []*user.User

	if err != nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, User: result, Error: err.Error()})
		return
	}

	defer resultData.Close()

	for resultData.Next() {
		var id int
		var name string
		var email string
		var password string
		var updated_at string
		var created_at string

		err = resultData.Scan(&id, &name, &email, &password, &updated_at, &created_at)

		if err != nil {
			json.NewEncoder(ress).Encode(RessponseData{Success: false, User: result, Error: err.Error()})
			return
		}

		result = append(result, &user.User{ID: id, Name: name, Email: email, Password: password})
	}

	if len(result) >= 1 {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, User: result})
		return
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, User: result})
		return
	}
}

func GetUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)

	type RessponseData struct {
		Success bool
		User    user.User
		Error   string
	}

	idUser, err := strconv.Atoi(params["id"])

	if err != nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, User: user.User{}, Error: "User tidak ditemukan!"})
		return
	}

	resultData, err := db.Db().Query("SELECT * FROM `crud-test`.users WHERE id = ?", idUser)

	defer resultData.Close()

	resultData.Next()

	var id int
	var name string
	var email string
	var password string

	resultData.Scan(&id, &name, &email, &password)

	responseData := user.User{ID: id, Name: name, Email: email, Password: password}

	if responseData.ID != 0 {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, User: responseData, Error: ""})
		return
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, User: responseData, Error: "User tidak ditemukan!"})
		return
	}
}

func CreateUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	var user user.User

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

	_, err = db.Db().Query("INSERT INTO `crud-test`.users VALUES(0,?,?,?)", name, email, password)

	if err == nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, Message: "Berhasil menambah user!"})
		return
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, Message: "Gagal menambah user!"})
		return
	}
}

func UpdateUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	var user user.User

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

	_, err = db.Db().Query("UPDATE `crud-test`.users SET name = ?, email = ?, password = ? WHERE id = ?", name, email, password, idUser)

	if err == nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, Message: "Berhasil update data user!"})
		return
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, Message: "Gagal update data user!"})
		return
	}
}

func DeleteUser(ress http.ResponseWriter, router *http.Request) {
	ress.Header().Set("Content-Type", "application/json")
	var user user.User

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

	_, err = db.Db().Query("DELETE FROM `crud-test`.users WHERE id = ?", idUser)

	if err == nil {
		json.NewEncoder(ress).Encode(RessponseData{Success: true, Message: "Berhasil delete data user!"})
		return
	} else {
		json.NewEncoder(ress).Encode(RessponseData{Success: false, Message: "Gagal delete data user!"})
		return
	}
}
