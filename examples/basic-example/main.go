package main

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/belmegatron/gofair"
	"github.com/belmegatron/gofair/config"
	"github.com/belmegatron/gofair/streaming"
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

	client, err := gofair.NewClient(cfg, streaming.IntegrationEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Logging to the Exchange API will allow us to acquire a SessionToken
	_, err = client.Login()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Logged into Betfair Exchange.")

	filter := gofair.MarketFilter{}
	event_types, err := client.Betting.ListEventTypes(filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Events: %v", prettyPrint(event_types))
}
