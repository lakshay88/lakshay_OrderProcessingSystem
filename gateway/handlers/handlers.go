package handlers

import "github.com/lakshay88/lakshay_OrderProcessingSystem/database"

type Handler struct {
	db *database.Database
}

func NewHandlers(db database.Database) *Handler {
	return &Handler{db: &db}
}
