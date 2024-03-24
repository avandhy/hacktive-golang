package models

import "time"

type Table struct {
	ID           int `gorm:"primaryKey"`
	Number       string
	Capacity     int
	Status       string
	RestaurantID int
	Reservations []Reservation
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateTableRequest struct {
	Number       string
	Capacity     int
	RestaurantID int
}

type UpdateTableRequest struct {
	Number       string
	Capacity     int
	RestaurantID int
}
