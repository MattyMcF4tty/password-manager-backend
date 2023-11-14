package handlers

import (
	"Go-X-Supabase/utils"
	"net/http"
)

// Deletes user.
func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
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

	// Call the SignIn function from the utils package
	err = utils.DeleteUser(userData.ID)
	if err != nil {
		http.Error(w, "Error deleting user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the header and write the response for a successful operation
	w.WriteHeader(http.StatusOK)
	responseMessage := "User deleted successfully"
	_, writeErr := w.Write([]byte(responseMessage))
	if writeErr != nil {
		// Handle the error in writing response
		utils.LogError("Failed to write response:", writeErr)
	}
}
