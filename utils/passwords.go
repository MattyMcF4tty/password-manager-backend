package utils

import (
	"fmt"

	"Go-X-Supabase/supabase"
)

type Password struct {
	AppName  string
	Password string
}

func GetUserPasswords(userId string) ([]Password, error) {
	client := supabase.GetClient()

	var passwords []Password

	err := client.DB.From("Passwords").Select("*").Eq("userId", userId).Execute(&passwords)
	if err != nil {
		return nil, fmt.Errorf("error fetching passwords of user:%v. error: %v", userId, err)
	}

	fmt.Printf("Passwords: %v", passwords)
	fmt.Printf("Fetched passwords of user: %v", userId)

	return passwords, nil
}
