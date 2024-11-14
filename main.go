package main

import (
	"log"

	config "github.com/lakshay88/lakshay_OrderProcessingSystem/config"
	"github.com/lakshay88/lakshay_OrderProcessingSystem/database"
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
		}
	}
}

func main() {

}
