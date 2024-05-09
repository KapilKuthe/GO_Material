package database

import (
	"fmt"
	env "nfscGofiber/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabase(dsn string) (*gorm.DB, error) {
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Unable to connect to Database.\ncause: %v", err.Error())
		return nil, err
	}
	return db, nil
}

func InitializeDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort, env.DBSSLMode, env.DBTimezone)

	db, err := OpenDatabase(dsn)
	if err != nil {
		return nil
	}

	db.AutoMigrate()
	return db
}