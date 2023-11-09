package utils

import (
	"context"
	"fmt"

	"Go-X-Supabase/supabase"
)

type Password struct {
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
		return nil, fmt.Errorf("error fetching passwords of user:%v. error: %v", userId, err)
	}

	fmt.Println("Passwords:", passwords)
	fmt.Println("Fetched passwords of user:", userId)

	return &passwords, nil
}

// Creates a new password. Password is supposed to be encrypted on client. WILL NOT BE ENCYPRTED HERE.
func CreatePassword(userId string, appName string, password string) error {
	//Get supabase client
	client := supabase.GetClient()

	newPassword := Password{
		UserId:   userId,
		AppName:  appName,
		Password: password,
	}

	err := client.DB.From(passwordTable).Insert(newPassword).Execute(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create new password: %w", err)
	}

	fmt.Println("Created new password")
	return nil
}
