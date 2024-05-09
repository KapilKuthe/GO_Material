package main

import (
	env "nfscGofiber/environment"
	"nfscGofiber/msg"

	"github.com/gofiber/fiber/v2"
)


func init() {
	env.LoadEnvironmentVariables()
}

func main() {
    app := fiber.New()

	app.Post("/msg",msg.Notemsg)

    // Start the server
    app.Listen(":"+env.APP_PORT)
}
