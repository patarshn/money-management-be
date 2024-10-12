package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/patarshn/money-management-be/config"
	"github.com/patarshn/money-management-be/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Fail load env: %v", err)
	}

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Fail InitConfig: %v", err)
	}

	r, err := router.InitRouter(cfg)
	if err != nil {
		log.Fatalf("Fail InitRouter: %v", err)
	}

	router.Serve(cfg, r)
}
