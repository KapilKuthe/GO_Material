package main

import (
	"context"
	"log"
	"nfscGofiber/data"
	"nfscGofiber/database"
	env "nfscGofiber/environment"
	"nfscGofiber/msg"
	"time"

	"github.com/gofiber/fiber/v2"
)


func init() {
	env.LoadEnvironmentVariables()
}

func main() {
    app := fiber.New()

	// Establish a connection to PostgreSQL
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	go func() {
		dataRetriever := data.NewDataRetriever(db)
		for {
			data, err := dataRetriever.RetrieveData(context.Background())
			if err != nil {
				log.Printf("Error retrieving data: %v\n", err)
			} else {
				log.Println("Data retrieved successfully")
			}
			log.Println(data)
			time.Sleep(60 * time.Second)
		}
	}()


	app.Post("/addmsg",msg.Notemsg)
	// app.Post("/getmsg",msg.Notemsg)

    // Start the server
    app.Listen(":"+env.APP_PORT)
}
