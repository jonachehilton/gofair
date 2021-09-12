package streaming

import (
	"bufio"
	"crypto/tls"
)

type Client struct {
	stream *Stream
	scanner bufio.Scanner

	// TLS connection object to Streaming API endpoint
	conn *tls.Conn
	// TLS Certificate used to authenticate to Streaming API endpoint
	certs *tls.Certificate

	err chan error
	stop chan int
}

func (client *Client) connect() error {

	cfg := &tls.Config{Certificates: []tls.Certificate{*client.certs}}
	conn, err := tls.Dial("tcp", client.stream.endpoint, cfg)

	if err != nil {
		return err
	}

	c := bufio.NewReader(conn)
	buf, _, err := c.ReadLine()
	if err != nil {
		return err
	}

	connectionMessage := new(models.ConnectionMessage)
	err = connectionMessage.UnmarshalJSON(buf)
	if err != nil {
		return err
	}

	if connectionMessage.ConnectionID == "" {
		return &ConnectionError{}
	}

	client.connectionID = connectionMessage.ConnectionID
	client.conn = conn
	// This scanner allows us to keep reading bytes from the connection until we encounter "\r\n"
	client.scanner = *bufio.NewScanner(client.conn)
	client.scanner.Split(scanCRLF)

	return nil
}

func (l *StreamClient) authenticate() error {

	if l.conn == nil {
		return &NoConnectionError{}
	}

	authenticationMessage := new(models.AuthenticationMessage)
	authenticationMessage.SetID(l.uid)
	authenticationMessage.AppKey = l.client.Config.AppKey
	authenticationMessage.Session = l.client.Session.SessionToken

	b, err := authenticationMessage.MarshalJSON()
	if err != nil {
		return err
	}

	_, err = l.write(b)
	if err != nil {
		return err
	}

	buf, err := l.read()
	if err != nil {
		return err
	}

	statusMessage := new(models.StatusMessage)
	err = statusMessage.UnmarshalJSON(buf)
	if err != nil {
		return err
	}

	if statusMessage.StatusCode == "FAILURE" {

		l.log.WithFields(logrus.Fields{
			"errorCode":    statusMessage.ErrorCode,
			"errorMessage": statusMessage.ErrorMessage,
		}).Error("Failed to Authenticate")

		return &AuthenticationError{}
	}

	l.log.Debug("Authenticated")

	return nil
}


func (client *Client) write(b []byte) (int, error) {
	// Every message is in json & terminated with a line feed (CRLF)
	b = append(b, []byte{'\r', '\n'}...)
	return client.conn.Write(b)
}

func (client *Client) read() ([]byte, error) {

	client.scanner.Scan()

	if err := client.scanner.Err(); err != nil {
		return []byte{}, err
	}

	return client.scanner.Bytes(), nil
}

func (client *Client) readPump(errChan *chan error) {

	if client.conn == nil {
		err := new(NoConnectionError)
		*errChan <- err
		return
	}

	for {
		select {
		case <-client.stop:
			return
		default:
			buf, err := client.read()
			if err != nil {
				*errChan <- err
				return
			}



			stream.onData(op, buf)
		}
	}
}

func (client *Client) writePump(errChan *chan error) {
	for {
		select {

		case <-client.stop:
			return

		case marketSubscriptionMessage := <-client.marketSubscriptionRequest:
			b, err := marketSubscriptionMessage.MarshalJSON()
			if err != nil {
				*errChan <- err
				return
			}

			client.write(b)

		case orderSubscriptionMessage := <-client.orderSubscriptionRequest:
			b, err := orderSubscriptionMessage.MarshalJSON()
			if err != nil {
				*errChan <- err
				return
			}
			client.write(b)

		}
	}
}