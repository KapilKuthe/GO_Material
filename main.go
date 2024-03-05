package main

import (
	db "goNotification/database"
	"goNotification/service"

	"github.com/kataras/iris/v12"
)

func main() {

	db.InitializeDB() // Initialize database connection
	// db.CloseDB()

	server := iris.Default()

	server.Post("/sendemail", service.Sendemail)

	server.Run(iris.Addr(":8080"))

}
