package utils

import (
	"context"
	"fmt"
	"log"

	"Go-X-Supabase/supabase"

	supa "github.com/nedpals/supabase-go"
)

type UserCreds struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	userId      string
}

// Handles user signin request and returns the users credentials.
func SignInUser(email string, password string) (*UserCreds, error) {
	//Get the supabase client.
	client := supabase.GetClient()

	ctx := context.Background()

	//Verify email and password.
	userData, err := client.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	//Create user object
	user := UserCreds{
		Email:       userData.User.Email,
		AccessToken: userData.AccessToken,
		userId:      userData.User.ID,
	}

	log.Println("Signed in user:", user.userId)
	return &user, nil
}

// Signs up user if email is not already taken.
func SignUpUser(email string, password string) (*supa.User, error) {
	//Get supabase client
	client := supabase.GetClient()

	ctx := context.Background()

	//Create user if email is not already taken
	userData, err := client.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Println("Error signing in up user:", err)
		return nil, fmt.Errorf("failed to sign up new user")
	}

	log.Println("Signed up new user:", userData.ID)
	return userData, nil
}

// SignOutUser signs out a user using their access token.
// It returns an error if the sign out process fails.
func SignOutUser(accessToken string) error {
	// Get supabase client
	client := supabase.GetClient()

	// Create a background context
	ctx := context.Background()

	// Attempt to sign out the user
	err := client.Auth.SignOut(ctx, accessToken)
	if err != nil {
		// Log the error and return it
		log.Printf("Error signing out user: %v", err)
		return fmt.Errorf("failed to sign out user")
	}

	return nil
}

// Takes an access_token an verifies it.
func VerifyAccessToken(accessToken string) (*supa.User, error) {
	//Get supabase client
	client := supabase.GetClient()

	user, err := client.Auth.User(context.Background(), accessToken)
	if err != nil {
		LogError("Failed to verify access token:", err)
		return nil, fmt.Errorf("failed to verify authentication")
	}

	LogSuccess("Verified access token of user: " + user.ID)
	return user, nil
}
