package tcp_server

import (
	"net"
	"sync"
	"testing"
	"time"
)

func Test_accepting_new_client_callback(t *testing.T) {
	server := New("localhost:9999")

	var wg sync.WaitGroup
	wg.Add(3)

	var messageText string

	server.OnNewClient(func(c *Client) {
		wg.Done()
	})
	server.OnNewMessage(func(c *Client, message string) {
		wg.Done()
		messageText = message
	})
	server.OnClientConnectionClosed(func(c *Client, err error) {
		wg.Done()
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

	wg.Wait()

	if messageText != "Test message\n" {
		t.Error("received wrong message")
	}
}
