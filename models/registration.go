package models

import "time"

type User struct {
	ID               uint      `gorm:"primary_key" json:"-"`
	Username         string    `json:"username"`
	Email            string    `gorm:"unique" json:"email"`
	Password         string    `json:"password"`
	DOB              time.Time `gorm:"type:date" json:"date_of_birth,omitempty" form:"date_of_birth" time_format:"02-01-2006"`
	RegistrationDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"registration_date"`
	LastLogin        time.Time `json:"-"`
	IsActive         bool      `gorm:"default:true" json:"-"`
}

type JwtToken struct {
	ID           uint      `gorm:"primary_key" json:"-"`
	UserID       uint      `json:"-"`
	Token        string    `json:"token"`
	Expiration   time.Time `json:"expiration"`
	CreationTime time.Time `json:"creation_time"`
}
