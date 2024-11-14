package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/database"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/models"
)

// Requests
type GetCustomerByIDRequest struct {
	ID uint `json:"id" binding:"required,gt=0"`
}

type CreateCustomerRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type CreateOrderRequest struct {
	CustomerID uint   `json:"customer_id" binding:"required,gt=0"`
	ProductIDs []uint `json:"product_ids" binding:"required"`
}

type Handler struct {
	db database.Database
}

func NewHandlers(db database.Database) *Handler {
	return &Handler{db: db}
}

func (h Handler) GetAllCustomers() gin.HandlerFunc {
	return func(c *gin.Context) {
		customers, err := h.db.GetAllCustomers()
		if err != nil {
			RespondWithJSON(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		RespondWithJSON(c, http.StatusOK, customers)
	}
}

func (h *Handler) GetCustomerByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.Param("id")
		if customerID == "" {
			RespondWithJSON(c, http.StatusBadRequest, gin.H{
				"error": "Customer ID is required",
			})
			return
		}

		id, err := strconv.ParseUint(customerID, 10, 32)
		if err != nil {
			RespondWithJSON(c, http.StatusBadRequest, gin.H{
				"error": "Invalid customer ID format",
			})
			return
		}

		fmt.Println("================================================")
		fmt.Println(id)
		fmt.Println("================================================")

		customer, err := h.db.GetCustomerByID(uint(id))
		if err != nil {
			RespondWithJSON(c, http.StatusNotFound, gin.H{
				"error": "Customer not found",
			})
			return
		}

		// Respond with customer details
		RespondWithJSON(c, http.StatusOK, customer)
	}
}

func (h *Handler) CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateOrderRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				RespondWithJSON(c, http.StatusBadRequest, gin.H{
					"error":   "Invalid request payload",
					"details": validationErrs.Error(),
				})
			} else {
				RespondWithJSON(c, http.StatusBadRequest, gin.H{
					"error": "Invalid JSON format",
				})
			}
		}

		_, err := h.db.GetCustomerByID(req.CustomerID)
		if err != nil {
			RespondWithJSON(c, http.StatusNotFound, gin.H{
				"error": "Customer not found",
			})
			return
		}

		isUnfulfilled, err := h.db.IsPreviousOrderUnfulfilled(req.CustomerID)
		if err != nil {
			RespondWithJSON(c, http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Error checking previous order: %v", err),
			})
			return
		}
		if isUnfulfilled {
			RespondWithJSON(c, http.StatusBadRequest, gin.H{
				"error": "Cannot place a new order. Previous order is unfulfilled.",
			})
			return
		}

		products, err := h.db.GetProductsByIDs(req.ProductIDs)
		if err != nil {
			RespondWithJSON(c, http.StatusNotFound, gin.H{
				"error": "One or more products not found",
			})
			return
		}
		totalPrice := float64(0)
		for _, product := range products {
			totalPrice += product.Price
		}

		order := models.Order{
			CustomerID: req.CustomerID,
			Products:   products,
			TotalPrice: totalPrice,
		}

		if err := h.db.CreateOrder(&order); err != nil {
			RespondWithJSON(c, http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Respond with customer details
		RespondWithJSON(c, http.StatusOK, order)
	}
}

func (h *Handler) GetOrderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("id")
		if orderID == "" {
			RespondWithJSON(c, http.StatusBadRequest, gin.H{
				"error": "Order ID is required",
			})
			return
		}

		orderIDUint, err := strconv.ParseUint(orderID, 10, 32)
		if err != nil {
			RespondWithJSON(c, http.StatusBadRequest, gin.H{
				"error": "Invalid order ID format",
			})
			return
		}

		order, err := h.db.GetOrderByID(uint(orderIDUint))
		if err != nil {
			RespondWithJSON(c, http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Respond with the order details
		RespondWithJSON(c, http.StatusOK, order)
	}
}

func (h *Handler) CreateCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateCustomerRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			RespondWithJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Check if customer already exists by email
		existingCustomer, err := h.db.GetCustomerByEmail(req.Email)
		if err == nil && existingCustomer != nil {
			RespondWithJSON(c, http.StatusConflict, gin.H{"error": "Customer with this email already exists"})
			return
		}

		// Create a new customer
		customer := models.Customer{
			Name:  req.Name,
			Email: req.Email,
		}

		if err := h.db.CreateCustomer(&customer); err != nil {
			RespondWithJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})

			return
		}

		RespondWithJSON(c, http.StatusOK, customer)
	}
}

func (h *Handler) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateProductRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			RespondWithJSON(c, http.StatusInternalServerError, gin.H{"error": "Invalid request payload"})
			return
		}

		product := models.Product{
			Name:  req.Name,
			Price: req.Price,
		}

		if err := h.db.CreateProduct(&product); err != nil {
			RespondWithJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		RespondWithJSON(c, http.StatusOK, product)
	}
}

func RespondWithJSON(c *gin.Context, statusCode int, payload interface{}) {
	c.JSON(statusCode, payload)
}
