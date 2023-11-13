package utils

import (
	"fmt"

	"Go-X-Supabase/supabase"
)

// Type used to create new password
type Password struct {
	Id       int8   `json:"id"`
	UserId   string `json:"userId"`
	AppName  string `json:"appName"`
	Password string `json:"password"`
}

// Type used to create new password
type NewPassword struct {
	UserId   string `json:"userId"`
	AppName  string `json:"appName"`
	Password string `json:"password"`
}

// Setup
var passwordTable = "Passwords"

// Gets a users passwords. All are encrypted. User is supposed to decrypt on client.
func GetPasswords(userId string) (*[]Password, error) {
	client := supabase.GetClient()

	var passwords []Password

	err := client.DB.From(passwordTable).Select("*").Eq("userId", userId).Execute(&passwords)
	if err != nil {
		LogError("Could not fetch passwords of user "+userId, err)
		return nil, fmt.Errorf("error fetching passwords")
	}

	LogSuccess("Fetched passwords of user: " + userId)

	return &passwords, nil
}

// Creates a new password. Password is supposed to be encrypted on client. WILL NOT BE ENCYPRTED HERE.
func CreatePassword(userId string, appName string, password string) error {
	//Get supabase client
	client := supabase.GetClient()

	newPassword := NewPassword{
		UserId:   userId,
		AppName:  appName,
		Password: password,
	}

	var result []struct{} // Empty struct to satisfy the method signature

	err := client.DB.From(passwordTable).Insert(newPassword).Execute(&result)
	if err != nil {
		LogError("Failed to create new password for user"+userId, err)
		return fmt.Errorf("failed to create new password")
	}

	LogSuccess("Created new password for user " + userId)
	return nil
}

func DeletePassword(userId string, passwordId int8) error {
	client := supabase.GetClient()

	// Convert passwordId from int8 to string
	password, err := GetPassword(passwordId)
	if err != nil {
		LogError("Failed to delete password "+IntToString(int64(passwordId)), err)
		return fmt.Errorf("failed to delete password")
	}

	if userId != password.UserId {
		LogWarning("User " + userId + " tried to delete password " + IntToString(int64(passwordId)) + " which does not belong to user")
		return fmt.Errorf("You can't delete a password thats not yours")
	}

	// Request needs struct to finish
	var filler struct{}
	err = client.DB.From(passwordTable).Delete().Eq("id", IntToString(int64(passwordId))).Execute(filler)
	if err != nil {
		LogError("Failed to delete password", err)
		return fmt.Errorf("failed to delete password")
	}

	LogSuccess("Deleted password: " + IntToString(int64(passwordId)))
	return nil
}

func GetPassword(passwordId int8) (*Password, error) {
	// Get supabase client
	client := supabase.GetClient()

	passwordIdStr := IntToString(int64(passwordId))

	var passwords []Password // Use a slice of Password
	err := client.DB.From(passwordTable).Select().Limit(1).Eq("id", passwordIdStr).Execute(&passwords)
	if err != nil {
		LogError("Failed to fetch password "+passwordIdStr, err)
		return nil, fmt.Errorf("failed to fetch password: %w", err)
	}

	if len(passwords) == 0 {
		// No password found with the given ID
		return nil, nil
	}

	LogSuccess("Fetched password " + passwordIdStr)
	return &passwords[0], nil // Return the first element of the slice
}
