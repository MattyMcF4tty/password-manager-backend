package main

import (
	//Default packages
	"log"
	"net/http"

	//Local packages
	"Go-X-Supabase/config"
	"Go-X-Supabase/routes"

	//External packages
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Set up routes for different functionalities
	routes.RegisterRoutes(router)

	// Get server configuration
	cfg := config.GetConfig()

	// Start server
	log.Printf("Server is running on http://%s\n", cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}
