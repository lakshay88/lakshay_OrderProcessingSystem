package models

import (
	"time"
)

type Customer struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Email     string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Orders    []Order `gorm:"foreignKey:CustomerID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
