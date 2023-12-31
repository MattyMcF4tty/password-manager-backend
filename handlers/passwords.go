package handlers

import (
	"Go-X-Supabase/utils"
	"encoding/json"
	"net/http"
)

func HandleGetPasswords(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameter
	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		http.Error(w, "Missing Authorization", http.StatusBadRequest)
		return
	}

	//Verify accessToken
	userData, err := utils.VerifyAccessToken(accessToken)
	if err != nil {
		http.Error(w, "Invalid Authorization", http.StatusUnauthorized)
		return
	}

	// Assume you have a global Supabase client or find a way to pass it here
	paswords, err := utils.GetPasswords(userData.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paswords)
}

func HandleCreatePassword(w http.ResponseWriter, r *http.Request) {
	//Get and verify access token from request header
	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		http.Error(w, "Missing Authorization", http.StatusBadRequest)
		return
	}
	userData, err := utils.VerifyAccessToken(accessToken)
	if err != nil {
		http.Error(w, "Invalid Authorization", http.StatusUnauthorized)
		return
	}

	//Get AppName and Password from body
	var req utils.Password
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.AppName == "" {
		http.Error(w, "Missing AppName parameter in request body", http.StatusBadRequest)
		return
	}
	if req.Password == "" {
		http.Error(w, "Missing Password parameter in request body", http.StatusBadRequest)
		return
	}

	err = utils.CreatePassword(userData.ID, req.AppName, req.Password)
	if err != nil {
		http.Error(w, "Error creating password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the header and write the response for a successful operation
	w.WriteHeader(http.StatusOK)
	responseMessage := "Password created successfully"
	_, writeErr := w.Write([]byte(responseMessage))
	if writeErr != nil {
		// Handle the error in writing response
		utils.LogError("Failed to write response:", writeErr)
	}
}

func HandleDeletePassword(w http.ResponseWriter, r *http.Request) {
	//Get and verify access token from request header
	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		http.Error(w, "Missing Authorization", http.StatusBadRequest)
		return
	}
	userData, err := utils.VerifyAccessToken(accessToken)
	if err != nil {
		http.Error(w, "Invalid Authorization", http.StatusUnauthorized)
		return
	}

	//Get PasswordId from body
	var req utils.Password
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Id == 0 {
		http.Error(w, "Missing id parameter in request body", http.StatusBadRequest)
		return
	}

	err = utils.DeletePassword(userData.ID, req.Id)
	if err != nil {
		http.Error(w, "Error deleting password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the header and write the response for a successful operation
	w.WriteHeader(http.StatusOK)
	responseMessage := "Password deleted successfully"
	_, writeErr := w.Write([]byte(responseMessage))
	if writeErr != nil {
		// Handle the error in writing response
		utils.LogError("Failed to write response:", writeErr)
	}
}

func HandleUpdatePassword(w http.ResponseWriter, r *http.Request) {
	//Get and verify access token from request header
	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		http.Error(w, "Missing Authorization", http.StatusBadRequest)
		return
	}
	userData, err := utils.VerifyAccessToken(accessToken)
	if err != nil {
		http.Error(w, "Invalid Authorization", http.StatusUnauthorized)
		return
	}

	//Get AppName and Password from body
	var req utils.Password
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.AppName == "" {
		http.Error(w, "Missing AppName parameter in request body", http.StatusBadRequest)
		return
	}
	if req.Password == "" {
		http.Error(w, "Missing Password parameter in request body", http.StatusBadRequest)
		return
	}
	if req.Id == 0 {
		http.Error(w, "Missing PasswordId parameter in request body", http.StatusBadRequest)
		return
	}

	req.UserId = userData.ID

	err = utils.UpdatePassword(req)
	if err != nil {
		http.Error(w, "Error updating password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the header and write the response for a successful operation
	w.WriteHeader(http.StatusOK)
	responseMessage := "Password updated successfully"
	_, writeErr := w.Write([]byte(responseMessage))
	if writeErr != nil {
		// Handle the error in writing response
		utils.LogError("Failed to write response:", writeErr)
	}
}
