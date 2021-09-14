package streaming

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/belmegatron/gofair/common"
)

func TestTLSConnection(t *testing.T) {

	// Arrange
	config, _ := common.LoadConfig("../config.json")
	cert, _ := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
	
	// Act
	conn, _ := NewTLSConnection(StreamIntegrationEndpoint, &cert)

	// Assert
	assert.NotNil(t, conn)
}