package main

import (
	"flag"
	"log"

	"github.com/jonachehilton/gofair"
	"github.com/jonachehilton/gofair/config"
	"github.com/jonachehilton/gofair/streaming"
	"github.com/jonachehilton/gofair/streaming/models"
)

func getRandomMarketID(client *gofair.Client) string {

	filter := gofair.MarketFilter{EventTypeIds: []string{"1"}}
	events, err := client.Betting.ListEvents(filter)
	if err != nil {
		log.Fatal(err)
	}

	if len(events) == 0 {
		log.Fatal("Unable to find any Events on the Exchange.")
	}

	eventID := events[0].Event.ID

	filter.EventIds = []string{eventID}
	marketProjection := []string{"RUNNER_DESCRIPTION", "EVENT"}

	marketCatalogues, err := client.Betting.ListMarketCatalogue(filter, marketProjection, "", 1)
	if err != nil {
		log.Fatal(err)
	}

	if len(marketCatalogues) == 0 {
		log.Fatalf("Unable to find any Markets for EventID %v.", eventID)
	}

	return marketCatalogues[0].MarketID
}

/*
Using the config_template.json, create a file called config.json in this directory and enter your Betfair Exchange API credentials.
You will also need to supply paths to your SSL cert and private key.
*/
func main() {

	configPath := flag.String("config", "config.json", "Path to config.json")

	// Load our config
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Loaded config.json")

	// Create a new Client object for interacting with the Exchange API. We are using the Streaming IntegrationEndpoint which is for testing.
	client, err := gofair.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created new gofair client.")

	// Logging to the Exchange API will allow us to acquire a SessionToken.
	_, err = client.Login()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Logged into Betfair Exchange.")

	marketID := getRandomMarketID(client)

	// Kick off our connection to the Betfair Exchange Stream API.
	err = client.Streaming.Start(streaming.IntegrationEndpoint, client.Session.SessionToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Started connection to Exchange Stream API.")

	filter := models.MarketFilter{MarketIds: []string{marketID}}
	dataFilter := models.MarketDataFilter{Fields: []string{string(gofair.PriceDataEnum.ExBestOffers), "EX_MARKET_DEF"}, LadderLevels: 1}
	client.Streaming.SubscribeToMarkets(&filter, &dataFilter)
	log.Printf("Sent subscription request for Market %v.", marketID)

	// Block until we receive a response to inform us that we have successfully subscribed.
	subscribedMarket := <-client.Streaming.Channels.MarketUpdate
	log.Printf("Subscribed to %v.", subscribedMarket.MarketID)
}
