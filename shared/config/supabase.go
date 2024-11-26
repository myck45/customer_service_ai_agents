package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/supabase-community/supabase-go"
)

func NewSupabaseClient() (*supabase.Client, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	supabaseClient, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		logrus.Fatalf("Error creating Supabase client: %v", err)
		return nil, fmt.Errorf("error creating Supabase client %v", err)
	}

	return supabaseClient, nil
}
