package database

import (
	"github.com/lakshay88/lakshay_OrderProcessingSystem/models"
)

type Database interface {
	Close() error
	GetAllCustomers() ([]models.Customer, error)
	GetCustomerByID(uint) (models.Customer, error)
	GetProductsByIDs([]uint) ([]models.Product, error)
	CreateOrder(*models.Order) error
	GetOrderByID(uint) (models.Order, error)
	IsPreviousOrderUnfulfilled(uint) (bool, error)
	CreateCustomer(customer *models.Customer) error
	GetCustomerByEmail(string) (*models.Customer, error)
	CreateProduct(*models.Product) error
}
