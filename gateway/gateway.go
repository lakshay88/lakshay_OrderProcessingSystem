package gateway

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/config"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/database"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/gateway/routers"
)

type Gateway struct{}

func NewGateway() *Gateway {
	return &Gateway{}
}

func (g *Gateway) RegisterGateWayService(cfg *config.AppConfig, db database.Database) error {
	log.Println("Setting routers")

	r := gin.Default()
	apiRouter := routers.NewRouter()
	apiRouter.RegisterRoutes(r, db)

	// http server config
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.ServerConfig.ServerPort),
		Handler:      r,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 0,
		IdleTimeout:  60 * time.Second,
	}

	// channel to listen for OS signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// running in new goroutine to make main funtion free
	go func() {
		log.Println("Server started successfully on port", cfg.ServerConfig.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-signalChan
	log.Println("starting shutdown...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server Shutdown Error: %v", err)
	}
	log.Println("Server gracefully shut down")

	return nil
}
