package streaming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventHandler(t *testing.T) {
	// Arrange
	channels := newStreamChannels()
	marketCache := make(CachedMarkets)
	orderCache := make(CachedOrders)

	// Act
	handler := newEventHandler(channels, &marketCache, &orderCache)

	// Assert
	assert.NotNil(t, handler.Markets)
	assert.NotNil(t, handler.Orders)
}
