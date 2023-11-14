package utils

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func DeleteUser(userId string) error {
	// Load env
	if err := godotenv.Load(); err != nil {
		LogFatal("Error loading .env file when trying to delete supabase user", err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseServiceKey := os.Getenv("SUPABASE_SERVICE_KEY")

	// Create http client
	httpClient := &http.Client{}

	// Create manual request to supabase, as package does not support deleting a user at this moment
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf(supabaseUrl+"/auth/v1/admin/users/%s", userId),
		nil,
	)
	if err != nil {
		LogError("Failed to create delete supabase user http request", err)
		return fmt.Errorf("Supabase user deletion failed")
	}

	// Set auth headers
	req.Header.Set("apikey", supabaseServiceKey)
	req.Header.Set("Authorization", "Bearer "+supabaseServiceKey)
	req.Header.Set("Content-Type", "application/json")

	// Send http request
	res, err := httpClient.Do(req)
	if err != nil {
		LogError("Failed send delete supabase user http request", err)
		return fmt.Errorf("Supabase user deletion failed")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		LogError("Failed to delete supabase user", fmt.Errorf(res.Status))
		return fmt.Errorf("Supabase user deletion failed")
	}

	LogSuccess("User " + userId + " deleted successfully")
	return nil
}
