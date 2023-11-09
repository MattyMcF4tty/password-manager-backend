package handlers

import (
	"Go-X-Supabase/utils"
	"encoding/json"
	"net/http"
)

// Define a struct to match the expected JSON structure of the request body
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleSignInUser(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body into the SignInRequest struct
	var req SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Check for missing parameters
	if req.Email == "" {
		http.Error(w, "Missing email parameter in request body", http.StatusBadRequest)
		return
	}
	if req.Password == "" {
		http.Error(w, "Missing password parameter in request body", http.StatusBadRequest)
		return
	}

	// Call the SignIn function from the utils package
	user, err := utils.SignInUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Error signing in: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Signs up user if email is not already taken.
func HandleSignUpUser(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body into the SignInRequest struct
	var req SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Check for missing parameters
	if req.Email == "" {
		http.Error(w, "Missing email parameter in request body", http.StatusBadRequest)
		return
	}
	if req.Password == "" {
		http.Error(w, "Missing password parameter in request body", http.StatusBadRequest)
		return
	}

	// Call the SignIn function from the utils package
	user, err := utils.SignUpUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Error signing up: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
