package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/database"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/gateway/handlers"
)

type Routers struct{}

func NewRouter() *Routers {
	return &Routers{}
}

func (r *Routers) RegisterRoutes(router *gin.Engine, db database.Database) {

	handlersInstance := handlers.NewHandlers(db)

	// Customer routes
	router.POST("/api/customers", handlersInstance.CreateCustomer())
	router.GET("/api/customers", handlersInstance.GetAllCustomers())
	router.GET("/api/customers/:id", handlersInstance.GetCustomerByID())

	// Order routes
	router.POST("/api/orders", handlersInstance.CreateOrder())
	router.GET("/api/orders/:id", handlersInstance.GetOrderByID())

	// Product
	router.POST("/api/products", handlersInstance.CreateProduct())
}
