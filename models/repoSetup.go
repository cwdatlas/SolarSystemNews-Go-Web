package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	fmt.Print(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	err = database.AutoMigrate(&Article{})
	if err != nil {
		return
	}
	DB = database
}
