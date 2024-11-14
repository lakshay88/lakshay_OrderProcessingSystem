package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/config"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/database"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/gateway/handlers"
)

type Routers struct{}

func NewRouter() *Routers {
	return &Routers{}
}

func (r *Routers) RegisterRoutes(router *gin.Engine, cfg *config.AppConfig, db database.Database) {

	handlersInstance := handlers.NewHandlers(db)

}
