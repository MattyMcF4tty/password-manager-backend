package utils

import (
	"Go-X-Supabase/supabase"

	"fmt"
)

type User struct {
	ID    int8   `json:"id"`
	Email string `json:"email"`
	// ... include other fields as necessary
}

func GetUserById(id string) (*User, error) {
	//Get client
	client := supabase.GetClient()
	var users []User

	fmt.Println(id)
	// Execute the query and unmarshal the JSON response into the users slice
	err := client.DB.From("Users").Select("*").Eq("id", id).Execute(&users)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by email: %v", err)
	}
	fmt.Println(users)

	// Check if we got exactly one user back
	if len(users) == 1 {
		return &users[0], nil // Return the first user
	} else if len(users) > 1 {
		return nil, fmt.Errorf("multiple users found with the same email")
	} else {
		return nil, fmt.Errorf("no user found with the given email")
	}
}
