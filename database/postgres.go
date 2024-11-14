package database

import (
	"fmt"
	"log"
	"time"

	"github.com/lakshay88/lakshay_OrderProcessingSystem/config"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type PostgresDB struct {
	DB *gorm.DB
}

func ConnectDatabase(cfg *config.DatabaseConfig) (Database, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	log.Println("Establishing Database Connection")
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from GORM: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Run migrations
	if err := db.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database connected and migrated successfully")
	return &PostgresDB{DB: db}, nil
}

func (db *PostgresDB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from GORM: %w", err)
	}
	return sqlDB.Close()
}

func (db *PostgresDB) GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	if err := db.DB.Find(&customers).Error; err != nil {
		return nil, fmt.Errorf("not able to find customers: %w", err)
	}
	return customers, nil

}

func (db *PostgresDB) GetCustomerByID(id uint) (models.Customer, error) {
	var customer models.Customer
	err := db.DB.First(&customer, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Customer{}, fmt.Errorf("Customer not found with ID %d", id)
		}
		return models.Customer{}, fmt.Errorf("Error retrieving customer: %w", err)
	}
	return customer, nil
}

func (db *PostgresDB) GetProductsByIDs(productIDs []uint) ([]models.Product, error) {
	var products []models.Product
	if err := db.DB.Where("id IN ?", productIDs).Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to find products: %w", err)
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("no products found")
	}

	return products, nil
}

func (db *PostgresDB) CreateOrder(order *models.Order) error {
	if err := db.DB.Create(&order).Error; err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}
	return nil
}

func (db *PostgresDB) GetOrderByID(orderID uint) (models.Order, error) {
	var order models.Order
	if err := db.DB.Preload("Products").First(&order, orderID).Error; err != nil {
		return models.Order{}, fmt.Errorf("order not found: %w", err)
	}
	return order, nil
}

func (db *PostgresDB) IsPreviousOrderUnfulfilled(customerID uint) (bool, error) {
	var lastOrder models.Order
	err := db.DB.Where("customer_id = ? AND status = ?", customerID, "pending").
		Order("created_at desc").
		First(&lastOrder).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, fmt.Errorf("error fetching previous orders: %w", err)
	}

	if lastOrder.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (db *PostgresDB) GetCustomerByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	if err := db.DB.Where("email = ?", email).First(&customer).Error; err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}
	return &customer, nil
}

func (db *PostgresDB) CreateProduct(product *models.Product) error {
	if err := db.DB.Create(&product).Error; err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

func (db *PostgresDB) CreateCustomer(customer *models.Customer) error {
	return db.DB.Create(customer).Error
}
