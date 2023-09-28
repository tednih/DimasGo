package main

import (
	"DimasGo/config"
	"DimasGo/routes"
	"os"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	config.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + os.Getenv("PORT"))
}
