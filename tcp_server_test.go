package tcp_server

import (
	. "github.com/smartystreets/goconvey/convey"
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
	conn.Write([]byte("Test message\n"))

	time.Sleep(100 * time.Millisecond)

	conn.Close()

	// Wait for server
	time.Sleep(10 * time.Millisecond)

	server.Close()

	Convey("Messages should be equal", t, func() {
		So(messageText, ShouldEqual, "Test message\n")
	})
	Convey("It should receive new client callback", t, func() {
		So(newClient, ShouldEqual, true)
	})
	Convey("It should receive message callback", t, func() {
		So(messageReceived, ShouldEqual, true)
	})
	Convey("It should receive connection closed callback", t, func() {
		So(connectinClosed, ShouldEqual, true)
	})
}

func Test_accepting_new_client_callback_different_terminator(t *testing.T) {
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
	server.MessageTerminator = '\u0000'
	go server.Listen()

	// Wait for server
	// If test fails - increase this value
	time.Sleep(10 * time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		t.Fatal("Failed to connect to test server")
	}
	conn.Write([]byte("Test message\u0000"))

	time.Sleep(100 * time.Millisecond)

	conn.Close()

	// Wait for server
	time.Sleep(10 * time.Millisecond)

	server.Close()

	Convey("Messages should be equal", t, func() {
		So(messageText, ShouldEqual, "Test message\u0000")
	})
	Convey("It should receive new client callback", t, func() {
		So(newClient, ShouldEqual, true)
	})
	Convey("It should receive message callback", t, func() {
		So(messageReceived, ShouldEqual, true)
	})
	Convey("It should receive connection closed callback", t, func() {
		So(connectinClosed, ShouldEqual, true)
	})
}
