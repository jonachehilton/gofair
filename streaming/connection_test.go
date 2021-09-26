package streaming

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/belmegatron/gofair/config"
)

func TestTLSConnection(t *testing.T) {

	// Arrange
	cfg, _ := config.LoadConfig("../config.json")
	cert, _ := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	
	// Act
	conn, _ := newTLSConnection(IntegrationEndpoint, &cert)

	// Assert
	assert.NotNil(t, conn)
}