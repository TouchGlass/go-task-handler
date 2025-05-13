package db

import (
	"BDproj/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=mypassword dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err := DB.AutoMigrate(&service.Task{}); err != nil {
		log.Fatalf("Couldn't migrate: %v", err)
	}

	return DB, nil
}
