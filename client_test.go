package gofair

import (
	"testing"

	"github.com/belmegatron/gofair/config"
	"github.com/belmegatron/gofair/streaming"
	"github.com/stretchr/testify/assert"
)

func TestOrders(t *testing.T) {
	// Arrange
	cfg, _ := config.LoadConfig("config.json")
	client, _ := NewClient(cfg, streaming.IntegrationEndpoint)
	client.Login()
	client.Streaming.Start(client.Session.SessionToken)

	// Act
	client.Streaming.SubscribeToOrders()
	orderUpdate := <- client.Streaming.Channels.OrderUpdate
	
	// Assert
	assert.NotNil(t, orderUpdate)
}