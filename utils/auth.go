package utils

import (
	"context"
	"fmt"

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

	fmt.Println("Signed in user email:", user.Email)
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
		return nil, err
	}

	return userData, nil
}

// Takes an access_token an verifies it.
func VerifyAccessToken(accessToken string) (*supa.User, error) {
	//Get supabase client
	client := supabase.GetClient()

	user, err := client.Auth.User(context.Background(), accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token:%w", err)
	}

	fmt.Println("Verified access token of user:", user.ID)
	return user, nil
}
