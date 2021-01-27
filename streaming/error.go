package streaming

import (
	"fmt"
)

type NoConnectionError struct{}

func (err *NoConnectionError) Error() string {
	return fmt.Sprintf("No stream connection exists.")
}

type AuthenticationError struct{}

func (err *AuthenticationError) Error() string {
	return fmt.Sprintf("Failed to authenticate")
}

type ConnectionError struct{}

func (err *ConnectionError) Error() string {
	return fmt.Sprintf("Failed to connect.")
}
