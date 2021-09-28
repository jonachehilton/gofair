package main

import (
	"log"

	"github.com/belmegatron/gofair"
	"github.com/belmegatron/gofair/config"
	"github.com/belmegatron/gofair/streaming"
)

/*
	Using the config_template.json, create a file called config.json in this directory and enter your Betfair Exchange API credentials.
	You will also need to supply paths to your SSL cert and private key.
*/
func main() {

	// Load our local config
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Create a gofair client which allows us to interact with the Betfair Exchange API
	client, err := gofair.NewClient(cfg, streaming.IntegrationEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// This is just an example of a listener that will process responses from the Betfair Stream API
	go func(client *gofair.Client) {

		for {
			select {
			case err := <-client.Streaming.Channels.Err:
				log.Fatal(err)
			case marketUpdate := <-client.Streaming.Channels.MarketUpdate:
				log.Printf("Received a market update for MarketID: %v", marketUpdate.MarketID)
			case orderUpdate := <-client.Streaming.Channels.OrderUpdate:
				log.Printf("Received an order update for MarketID: %v", orderUpdate.MarketID)
			case marketSubscription := <-client.Streaming.Channels.MarketSubscriptionResponse:
				log.Printf("Subscribed to Markets: %v", marketSubscription.SubscribedMarketIDs)
			}
		}
	}(client)

	// Logging to the Exchange API will allow us to acquire a SessionToken
	_, err = client.Login()
	if err != nil {
		log.Fatal(err)
	}

	// Kick off our connection to the Betfair Exchange Stream API
	err = client.Streaming.Start(client.Session.SessionToken)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Obtain a MarketID from somewhere that we can subscribe to via the Stream API

}
