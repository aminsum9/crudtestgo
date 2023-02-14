package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() {
	// Replace the DB_USER and DB_PASSWORD with your database credentials
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Test the connection
	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the database!")
}
