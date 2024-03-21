package database

import (
	"assign-2/pkg/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=user dbname=golang port=5432 sslmode=disable "

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// migrate database
	db.AutoMigrate(&models.Order{}, &models.Item{})

	fmt.Println("Successfully connected!")

	return db
}
