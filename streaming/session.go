package streaming

import (
	"bufio"
	"crypto/tls"

	"github.com/belmegatron/gofair/streaming/models"
)

const (
	// Error Codes for Authenticating
	failure = "FAILURE"
)

type channels struct {
	err  chan error
	stop chan int
}

type Session struct {
	conn         *TLSConnection
	appKey       string
	sessionToken string
	scanner      *bufio.Scanner
	channels     channels
}

func NewSession(destination string, certs *tls.Certificate, appKey string, sessionToken string) (*Session, error) {
	session := new(Session)
	conn, err := NewTLSConnection(destination, certs)
	if err != nil {
		return nil, err
	}

	session.conn = conn
	err = session.authenticate(appKey, sessionToken)

	return session, nil
}

func (session *Session) authenticate(appKey string, sessionToken string) error {

	if session.conn == nil {
		return &NoConnectionError{}
	}

	authenticationMessage := &models.AuthenticationMessage{AppKey: appKey, Session: sessionToken}
	authenticationMessage.SetID(session.conn.ID)

	b, err := authenticationMessage.MarshalJSON()
	if err != nil {
		return err
	}

	_, err = session.write(b)
	if err != nil {
		return err
	}

	buf, err := session.read()
	if err != nil {
		return err
	}

	statusMessage := new(models.StatusMessage)
	err = statusMessage.UnmarshalJSON(buf)
	if err != nil {
		return err
	}

	if statusMessage.StatusCode == failure {
		return &AuthenticationError{}
	}

	return nil
}

func (session *Session) write(b []byte) (int, error) {
	// Every message is in json & terminated with a line feed (CRLF)
	b = append(b, []byte{'\r', '\n'}...)
	return session.conn.Write(b)
}

func (session *Session) read() ([]byte, error) {

	session.scanner.Scan()

	if err := session.scanner.Err(); err != nil {
		return []byte{}, err
	}

	return session.scanner.Bytes(), nil
}

func (session *Session) readPump() {

	if session.conn == nil {
		err := new(NoConnectionError)
		session.channels.err <- err
		return
	}

	for {
		select {
		case <-session.channels.stop:
			return
		default:
			buf, err := session.read()
			if err != nil {
				session.channels.err <- err
				return
			}

			stream.onData(op, buf)
		}
	}
}

func (session *Session) writePump() {
	for {
		select {

		case <-session.channels.stop:
			return

		case marketSubscriptionMessage := <-session.marketSubscriptionRequest:
			b, err := marketSubscriptionMessage.MarshalJSON()
			if err != nil {
				session.channels.err <- err
				return
			}

			session.write(b)

		case orderSubscriptionMessage := <-session.orderSubscriptionRequest:

			b, err := orderSubscriptionMessage.MarshalJSON()
			if err != nil {
				session.channels.err <- err
				return
			}

			session.write(b)
		}
	}
}
