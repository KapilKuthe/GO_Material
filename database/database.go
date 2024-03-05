package database

import (
	"fmt"
	"goNotification/model"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitializeDB() {
	username := "postgres"
	password := "Pass#1230"
	dbName := "test"
	host := "localhost"
	port := 5432

	escapedPassword := url.QueryEscape(password)
	dns := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", username, escapedPassword, host, port, dbName)
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failure! DB connection not established...")
	}
	fmt.Println("Connected...")
	// db.AutoMigrate() // Assuming you have a User model defined in the same package
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sqlDB.Close()
}

func GetTemplate(comtmpl model.MSCommunication) (model.MSCommunication, error) {
	var nwtempl model.MSCommunication

	result := db.Where("tmpl_id = ?", comtmpl.TmplID).First(&nwtempl)
	// result := db.First(&nwtempl, comtmpl.TmplID)
	if result.Error != nil {
		// fmt.Println(result.Error)
		return model.MSCommunication{}, result.Error
	}
	return nwtempl, nil
}
