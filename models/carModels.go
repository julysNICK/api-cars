package models

import "time"

type Car struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Make      string    `json:"make"`
	Model     string    `json:"model"`
	Year      string    `json:"year"`
	Is_Sold   bool      `json:"isSold"`
	Color     string    `json:"color"`
	Price     int       `json:"price"`
	User_Id   uint      `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
