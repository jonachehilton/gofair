package streaming

type NoConnectionError struct{}

func (err *NoConnectionError) Error() string {
	return "No stream connection exists"
}

type AuthenticationError struct{}

func (err *AuthenticationError) Error() string {
	return "Failed to authenticate"
}

type ConnectionError struct{}

func (err *ConnectionError) Error() string {
	return "Failed to connect"
}

type EndpointError struct{}

func (err *EndpointError) Error() string {
	return "Invalid stream endpoint"
}
