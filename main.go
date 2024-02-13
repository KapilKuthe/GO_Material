package main

import (
	"goLogin/database"
	"goLogin/routes"
)

func main() {
	// fmt.Println("hello world!")

	//? DB connection
	database.InitializeDB()

	routes.InitializeRoutes()

	database.CloseDB()
}
