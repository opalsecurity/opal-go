package main

import (
	"context"
	"fmt"
	"os"

	opal "github.com/opalsecurity/opal-go"
)

func main() {
	// Create a new API client configuration
	configuration := opal.NewConfiguration()
	// Uncomment the following lines to set a specific server URL
	// configuration.Servers = opal.ServerConfigurations{
	// 	{
	// 		URL: "https://api.opal.com/v1",
	// 	},
	// }

	// Set API token from environment variable
	apiToken := os.Getenv("OPAL_API_TOKEN")
	if apiToken == "" {
		fmt.Println("Error: OPAL_API_TOKEN environment variable is not set")
		os.Exit(1)
	}

	// Create a context with the API token for authentication
	ctx := context.WithValue(context.Background(), opal.ContextAccessToken, apiToken)

	// Initialize the API client
	client := opal.NewAPIClient(configuration)

	// Example: Get users
	resources, _, err := client.ResourcesAPI.GetResources(ctx).Execute()
	if err != nil {
		fmt.Printf("Error fetching resources: %v\n", err)
		os.Exit(1)
	}

	// Print the number of resources fetched
	fmt.Printf("Fetched %d resources\n", len(resources.GetResults()))
}
