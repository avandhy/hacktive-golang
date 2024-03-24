package models

import "time"

type Restaurant struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Location    string
	Description string
	Tables      []Table
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
