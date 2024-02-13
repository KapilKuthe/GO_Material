package database

import (
	"fmt"
	"goLogin/models"
	"goLogin/utility"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitializeDB() {
	username := "postgres"
	password := "Pass#1230"
	dbName := "myApp"
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
	db.AutoMigrate(&models.Customer{}, &models.User{}, &models.JwtToken{}) // Assuming you have a User model defined in the same package
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sqlDB.Close()
}

// ? create new customer
func CreateCustomer(customer models.Customer) (models.Customer, error) {
	result := db.Create(&customer)
	if result.Error != nil {
		// fmt.Println(result.Error)
		return models.Customer{}, result.Error
	}
	return customer, nil
}

// ? get all customer
func GetAllCustomer() ([]models.Customer, error) {
	var customer []models.Customer
	result := db.Find(&customer)
	if result.Error != nil {
		// fmt.Println(result.Error)
		return nil, result.Error
	}
	return customer, nil
}

// ? update customer
func UpdateCustomer(customer models.Customer) (models.Customer, error) {
	var existingCustomer models.Customer

	//* fetching  detail form Id
	result := db.First(&existingCustomer, customer.Id)
	if result.Error != nil {
		fmt.Println(result.Error)
		return models.Customer{}, result.Error
	}

	//* updating the new detail
	db.Model(&existingCustomer).Updates(customer)

	//* Fetch the updated record
	updatedCustomer := models.Customer{}
	db.First(&updatedCustomer, customer.Id)

	return updatedCustomer, nil
}

// ? Delete customer
func DeleteCustomer(customerId uint64) error {
	result := db.Delete(&models.Customer{}, customerId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

// ? create new user
func CreateUser(user models.User) (models.User, error) {
	password, err := utility.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.Password = password
	result := db.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

// ? get user
func GetUser(user models.User) (models.User, error) {
	var existingUser models.User
	result := db.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		// fmt.Println(result.Error)
		return models.User{}, result.Error
	}
	return existingUser, nil
}
