package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/a-poor/usps"
	_ "github.com/joho/godotenv/autoload"
)

// Load the environment variables from .env file
var APIUserID = os.Getenv("API_USER")

func getAddresses() []usps.Address {
	if len(os.Args) < 2 {
		panic("Missing argument: []usps.Address")
	}

	raw := os.Args[1]

	var addresses []usps.Address
	if err := json.Unmarshal([]byte(raw), &addresses); err != nil {
		panic(err)
	}

	return addresses
}

func main() {
	if APIUserID == "" {
		fmt.Println("API_USER not set")
		os.Exit(1)
	}

	// Load the addresses from the command line
	addrs := getAddresses()

	// Create the client
	c := usps.New(APIUserID)

	// Create a context
	ctx := context.Background()

	// Send the request
	req := &usps.ValidateAddressRequest{
		Addresses: addrs,
	}
	res, err := c.ValidateAddress(ctx, req)
	if err != nil {
		panic(err)
	}

	// Print the response (as JSON)
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
