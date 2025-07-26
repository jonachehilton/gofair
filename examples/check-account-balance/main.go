package main

import (
	"flag"
	"log"

	"github.com/jonachehilton/gofair"
	"github.com/jonachehilton/gofair/config"
)

func main() {
	configPath := flag.String("config", "config.json", "Path to config.json")

	// Load our config
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Loaded config.json")

	// Create a new Client object which we will use for interacting with the Exchange API
	client, err := gofair.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Login to the Exchange API
	_, err = client.Login()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Logged into Betfair Exchange.")

	// Check what our available balance is on the Exchange
	response, err := client.Account.GetAccountFunds()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Available Funds: %v", response.AvailableToBetBalance)
}
