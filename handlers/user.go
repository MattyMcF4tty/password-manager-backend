package handlers

import (
	"Go-X-Supabase/utils"
	"encoding/json"
	"net/http"
)

func HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}

	// Assume you have a global Supabase client or find a way to pass it here
	user, err := utils.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
