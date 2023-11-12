package routes

import (
	"Go-X-Supabase/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	//Auth
	router.HandleFunc("/auth/signin", handlers.HandleSignInUser).Methods("POST")
	router.HandleFunc("/auth/signup", handlers.HandleSignUpUser).Methods("POST")
	//	router.HandleFunc("/auth/signout", handlers.HandleGetUserById).Methods("GET")

	//Passwords
	router.HandleFunc("/passwords/get", handlers.HandleGetPasswords).Methods("GET")
	router.HandleFunc("/passwords/create", handlers.HandleCreatePassword).Methods("POST")
	router.HandleFunc("/passwords/delete", handlers.HandleDeletePassword).Methods("DELETE")
}
