package models

import "time"

type Order struct {
	ID         int       `json:"id" gorm:"primarykey"`
	Name       string    `json:"Name"`
	Ordered_at time.Time `json:"Ordered_at"`
	Items      []Item    `json:"Items" gorm:"ForeignKey:OrderID"`
}

type CreateOrderRequest struct {
	Name  string `json:"name" binding:"required"`
	Items []Item `json:"items" binding:"required"`
}

type UpdateOrderRequest struct {
	Name  string `json:"name" binding:"required"`
	Items []Item `json:"items" binding:"required"`
}
