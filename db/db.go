package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-test")

	if err != nil {
		panic(err.Error())
	}

	Database = db

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the database!")
}
