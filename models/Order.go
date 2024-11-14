package models

import "time"

type Order struct {
	ID         uint      `gorm:"primaryKey"`
	CustomerID uint      `gorm:"not null"`
	Customer   Customer  `gorm:"foreignKey:CustomerID"`
	Status     string    `gorm:"type:varchar(20);default:'unfulfilled'"`
	TotalPrice float64   `gorm:"not null"`
	Products   []Product `gorm:"many2many:order_products;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
