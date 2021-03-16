package tcp_server

import (
	"net"
	"testing"
	"time"
)

func buildTestServer() *server {
	return New("localhost:9999")
}

func Test_accepting_new_client_callback(t *testing.T) {
	server := buildTestServer()

	var messageReceived bool
	var messageText string
	var newClient bool
	var connectinClosed bool

	server.OnNewClient(func(c *Client) {
		newClient = true
	})
	server.OnNewMessage(func(c *Client, message string) {
		messageReceived = true
		messageText = message
	})
	server.OnClientConnectionClosed(func(c *Client, err error) {
		connectinClosed = true
	})
	go server.Listen()

	// Wait for server
	// If test fails - increase this value
	time.Sleep(10 * time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		t.Fatal("Failed to connect to test server")
	}
	_, err = conn.Write([]byte("Test message\n"))
	if err != nil {
		t.Fatal("Failed to send test message.")
	}
	conn.Close()

	// Wait for server
	time.Sleep(10 * time.Millisecond)

	if messageText != "Test message\n" {
		t.Error("received wrong message")
	}

	if newClient != true {
		t.Error("OnNewClient did not received call")
	}

	if messageReceived != true {
		t.Error("the message was not received")
	}

	if connectinClosed != true {
		t.Error("connection was not closed")
	}
}
