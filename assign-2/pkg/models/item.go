package models

type Item struct {
	ID          int    `json:"ID" gorm:"primarykey"`
	Item_code   string `json:"item_code" `
	Description string `json:"description" `
	Quantity    int    `json:"quantity" `
	OrderID     int    `json:"OrderID"`
}
