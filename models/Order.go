package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint      `json:"customer_id" gorm:"not null"`
	Products   []Product `json:"products" gorm:"many2many:order_products;"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status" gorm:"default:'pending'"`
}
