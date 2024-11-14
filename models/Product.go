package models

import "time"

type Product struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Price     float64 `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
