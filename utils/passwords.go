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

	LogSuccess("Fetched passwords of user:" + userId)

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

	var passwords []Password
	err := client.DB.From(passwordTable).Insert(newPassword).Execute(&passwords)
	if err != nil {
		LogError("Failed to create new password for user"+userId, err)
		return fmt.Errorf("failed to create new password")
	}

	LogSuccess("Created new password for user " + userId)
	return nil
}
