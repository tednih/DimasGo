package main

import (
	"DimasGo/config"
	"DimasGo/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.InitDB()

	r := routes.SetupRouter()

	r.Run(":8080")
}
