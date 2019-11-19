package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "")
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(200)
	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("db init")
}

func BeginTx() (*sql.Tx, error) {
	return db.Begin()
}

func checkErr(err error) {
	if err != nil {
		println(err.Error())
		log.Println(err.Error())
	}
}
