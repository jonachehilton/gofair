package gofair

import (
	"log"
	"testing"

	"github.com/belmegatron/gofair/config"
	"github.com/belmegatron/gofair/streaming"
	"github.com/stretchr/testify/assert"
)

func TestOrders(t *testing.T) {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	client, err := NewClient(cfg, streaming.IntegrationEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Login()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Streaming.Start(client.Session.SessionToken)
	if err != nil  {
		log.Fatal(err)
	}

	// Act
	client.Streaming.SubscribeToOrders()
	orderUpdate := <-client.Streaming.Channels.OrderUpdate

	// Assert
	assert.NotNil(t, orderUpdate)
}
