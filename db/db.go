package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func TestDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the database!")
}

func Db() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-test")

	if err != nil {
		panic(err.Error())
	}

	return db
}
