package models

type Customer struct {
	Id     uint64 `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Dob    string `json:"dob"`
	Mobile uint64 `json:"mobile"`
	Email  string `json:"email"`
}
