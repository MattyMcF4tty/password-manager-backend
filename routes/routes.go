package routes

import (
	"Go-X-Supabase/handlers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user/getbyid", handlers.HandleGetUserById).Methods("GET")

	router.HandleFunc("/user/getPasswords", handlers.HandleGetPasswords).Methods("GET")
	// ... register other user routes
}
