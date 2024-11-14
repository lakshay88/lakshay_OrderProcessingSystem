package main

import (
	"log"

	config "github.com/lakshay88/lakshay_OrderProcessingSystem/config"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/database"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/gateway"
)

var (
	cfg *config.AppConfig
	db  database.Database
)

func init() {
	log.Println("Starting Application started, fetching configurations")
	var err error
	cfg, err = config.LoadConfiguration("./config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		return
	}

	log.Println("Connection with  with db")
	switch cfg.Database.Driver {
	case "postgres":
		db, err = database.ConnectDatabase(&cfg.Database)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
			return
		}
	}
}

func main() {

	defer db.Close()

	gatewayInstance := gateway.NewGateway()

	err := gatewayInstance.RegisterGateWayService(cfg, db)
	if err != nil {
		log.Fatalf("Failed to initate gateway: %v", err)
		return
	}

}
