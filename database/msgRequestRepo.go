package database

import (
	"nfscGofiber/msg"

	"gorm.io/gorm"
)

// MsgRequestRepository handles CRUD operations for MsgRequest
type MsgRequestRepository struct {
	db *gorm.DB
}

// NewMsgRequestRepository creates a new MsgRequestRepository
func NewMsgRequestRepository(db *gorm.DB) *MsgRequestRepository {
	return &MsgRequestRepository{db: db}
}


// Create inserts a new MsgRequest record into the database
func (r *MsgRequestRepository) Create(msgReq *msg.MsgRequest) error {
	return r.db.Create(msgReq).Error
}

// GetAll retrieves all MsgRequest records from the database
func (r *MsgRequestRepository) GetAll() ([]msg.MsgRequest, error) {
	var msgReqs []msg.MsgRequest
	err := r.db.Find(&msgReqs).Error
	return msgReqs, err
}

// GetByID retrieves a MsgRequest record by its ID from the database
func (r *MsgRequestRepository) GetByID(id uint) (*msg.MsgRequest, error) {
	var msgReq msg.MsgRequest
	err := r.db.First(&msgReq, id).Error
	if err != nil {
		return nil, err
	}
	return &msgReq, nil
}

// Update updates a MsgRequest record in the database
func (r *MsgRequestRepository) Update(msgReq *msg.MsgRequest) error {
	return r.db.Save(msgReq).Error
}

// Delete deletes a MsgRequest record from the database
func (r *MsgRequestRepository) Delete(id uint) error {
	return r.db.Delete(&msg.MsgRequest{}, id).Error
}