// Local Supabase logic.
package supabase

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

// Returns the supabase client if initialized. If not then creates a new one and returns it.
var client *supa.Client

func GetClient() *supa.Client {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if client == nil {
		client = supa.CreateClient(supabaseUrl, supabaseKey)
		fmt.Println("Created new supabase client.")
	}

	return client
}
