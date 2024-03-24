package database

import (
	"final-project/pkg/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DataBaseInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=user dbname=golang port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Restaurant{}, &models.Table{}, &models.Reservation{})

	fmt.Println("Successfully connected!")

	return db

}
