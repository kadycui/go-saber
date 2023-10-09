package model

import "time"

type Book struct {
	Id        uint    `gorm:"primaryKey"`
	No        int64   `gorm:"not null"`
	Title     string  `gorm:"not null"`
	Author    string  `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Rating    float64 `gorm:"not null"`
	Date      string  `gorm:"not null"`
	Category  int     `gorm:"not null"`
	Publish   string  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
