package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	server   = "10.112.31.91"
	port     = 1433
	user     = "admin"
	password = "AdM!n@91"
	database = "test"
)

var db *sql.DB

func initDB() {
	var err error
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
}

// main.go (continuation)

func main() {
	initDB()
	defer db.Close()

}
