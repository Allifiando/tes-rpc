package models

import (
	"database/sql"
	"fmt"
	"os"

	// Register some standard stuff
	_ "github.com/go-sql-driver/mysql"
)

// Init ...
func Init() *sql.DB {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"))
	db, err := sql.Open("mysql", connectString)

	if err != nil {
		fmt.Println(err)
		panic("DB Connection Error")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic("DB Connection Error")
	}

	// Max connection
	db.SetMaxOpenConns(1)

	// fmt.Printf("DB Connect")
	return db

}
