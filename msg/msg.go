package msg

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	requests    []MsgRequest
	mutex       sync.Mutex
	activeCount int
)

func Notemsg(ctx *fiber.Ctx) error {
	var reqBody MsgRequest
	if err := ctx.BodyParser(&reqBody); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error(), "message": "Invalid request body"})
	}

	
	if err :=reqBody.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error(), "message": "Timestamp must be in future"})
	}

	// Save the message
	err := storemsg(reqBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error(), "message": "Error storing message: "})
	}
	
	log.Println("ActiveCount: ",activeCount)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// func Getmsg(ctx *fiber.Ctx) error {
	
// 	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": ctx.JSON()})
// }

func storemsg(reqBody MsgRequest) error {
	mutex.Lock()
	defer mutex.Unlock()

	requests = append(requests, reqBody)
	activeCount++
	// // Start a goroutine
	// go func() {
	// 	// Parse timestamp string into a time.Time object
	// 	timestamp := reqBody.Timestamp

	// 	// Calculate the duration
	// 	durationUntil := time.Until(timestamp)
	// 	time.Sleep(durationUntil)
	// 	log.Println("Message:", reqBody.Message)
	// 	removeMessage(reqBody)
	// }()
		
	
	// // Insert MsgRequest into database
	// if err := db.Create(&reqBody).Error; err != nil {
	// 	log.Printf("Error inserting MsgRequest: %v\n", err)
	// 	return fiber.NewError(fiber.StatusInternalServerError, "Error inserting data")
	// }
	return nil
}

func removeMessage(reqBody MsgRequest) {
	mutex.Lock()
	defer mutex.Unlock()

	for i, msg := range requests {
		if msg == reqBody {
			requests = append(requests[:i], requests[i+1:]...)
			activeCount--
			break
		}
	}
}

