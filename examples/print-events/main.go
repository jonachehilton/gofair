package main

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/jonachehilton/gofair"
	"github.com/jonachehilton/gofair/config"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

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

	// Fire off a request to list all available EventTypes
	filter := gofair.MarketFilter{}
	event_types, err := client.Betting.ListEventTypes(filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Events: %v", prettyPrint(event_types))
}
