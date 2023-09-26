// main.go
package main

import (
	"DimasGo/config"
	"DimasGo/routes"
)

func main() {
	// Inisialisasi koneksi database
	config.InitDB()

	// Setup router
	r := routes.SetupRouter()

	// Run server
	r.Run(":8080")
}
