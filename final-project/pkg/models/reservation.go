package models

import "time"

type Reservation struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Phone     int
	Date      string
	TableID   int
	UserID		uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateReservationRequest struct{
	Name		string
	Email		string
	Phone     	int
	Date      	string
	TableID   	int
	UserID		uint

}

type UpdateReservationRequest struct{
	Date      string
	TableID   int
}
