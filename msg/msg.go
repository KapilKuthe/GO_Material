package msg

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MsgRequest struct {
	Timestamp time.Time `json:"time"`
	Message   string    `json:"message"`
}

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

	// Validate message not empty
	if reqBody.Message == "" || strings.TrimSpace(reqBody.Message) == "" {
		// log.Println("Message is empty")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": errors.New("message not provided"), "message": "Message cannot be empty"})
	}

	// Validate timestamp not 0
	if reqBody.Timestamp.IsZero() {
		// log.Println("Timestamp is required")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": errors.New("timeStamp not provided"), "message": "Timestamp is required"})
	}

	// Check if timestamp is in future
	if reqBody.Timestamp.Before(time.Now()) {
		// log.Println("Timestamp is old")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": errors.New("old timeStamp"), "message": "Timestamp must be in future"})
	}

	// Save the message
	err := storemsg(reqBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error(), "message": "Error storing message: "})
	}
	
	log.Println("ActiveCount: ",activeCount)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func storemsg(reqBody MsgRequest) error {
	mutex.Lock()
	defer mutex.Unlock()

	requests = append(requests, reqBody)
	activeCount++
	// Start a goroutine
	go func() {
		// Parse timestamp string into a time.Time object
		timestamp := reqBody.Timestamp

		// Calculate the duration
		durationUntil := time.Until(timestamp)

		time.Sleep(durationUntil)
		log.Println("Message:", reqBody.Message)
		removeMessage(reqBody)
	}()
		
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

//we can add a db if data is needed to be stored but on what basis and why in db and not in redis catche