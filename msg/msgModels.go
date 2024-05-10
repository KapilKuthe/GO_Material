package msg

import (
	"time"
	"strings"
	"errors"
)

type MsgRequest struct {
	Id uint `gorm:"primary_key" json:"id"`
	Timestamp time.Time `json:"time"`
	Message   string    `json:"message"`
	IsActive bool `json:"isactive,omitempty" `
}


func (msgr *MsgRequest) Validate()error{
	// Validate message not empty
	if msgr.Message == "" || strings.TrimSpace(msgr.Message) == "" {
		// log.Println("Message is empty")
		return errors.New("message not provided")
	}

	// Validate timestamp not 0
	if msgr.Timestamp.IsZero() {
		// log.Println("Timestamp is required")
		return errors.New("timeStamp not provided")
	}

	// Check if timestamp is in future
	if msgr.Timestamp.Before(time.Now()) {
		// log.Println("Timestamp is old")
		return errors.New("old timeStamp")
	}
	return nil
}