package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name   string  `json:"name" gorm:"not null"`
	Email  string  `json:"email" gorm:"unique;not null"`
	Orders []Order `json:"orders" gorm:"foreignKey:CustomerID"`
}
