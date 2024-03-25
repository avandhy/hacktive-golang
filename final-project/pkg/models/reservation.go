package models

import "time"

type Reservation struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Phone     int
	Date      time.Time
	TableID   int
	UserID		uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateReservationRequest struct{
	Name		string
	Email		string
	Phone     	int
	Date      	time.Time
	TableID   	int
	UserID		uint

}

type UpdateReservationRequest struct{
	Date      time.Time
	TableID   int
}
