package tcpserver

import (
	"encoding/gob"
	"log"
	"net"
)

// CommunicationData is the data that is transmitted between client and server
type CommunicationData struct {
	Type string
	Data map[string]string
}

// Client holds info about connection
type Client struct {
	conn   net.Conn
	Server *server
}

// TCP server
type server struct {
	address                  string // Address to open connection: localhost:9999
	onNewClientCallback      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, response *CommunicationData)
}

// Read client data from channel
func (c *Client) listen() {
	decoder := gob.NewDecoder(c.conn)
	for {
		r := &CommunicationData{}
		err := decoder.Decode(r)
		if err != nil {
			c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}
		c.Server.onNewMessage(c, r)
	}
}

// Send text gob to client
func (c *Client) Send(message string, data map[string]string) error {
	encoder := gob.NewEncoder(c.conn)
	r := &CommunicationData{Type: message, Data: data}
	err := encoder.Encode(r)
	return err
}

// // Send bytes to client
// func (c *Client) SendBytes(b []byte) error {
// 	_, err := c.conn.Write(b)
// 	return err
// }

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// Called right after server starts listening new client
func (s *server) OnNewClient(callback func(c *Client)) {
	s.onNewClientCallback = callback
}

// Called right after connection closed
func (s *server) OnClientConnectionClosed(callback func(c *Client, err error)) {
	s.onClientConnectionClosed = callback
}

// Called when Client receives new message
func (s *server) OnNewMessage(callback func(c *Client, response *CommunicationData)) {
	s.onNewMessage = callback
}

// Start network server
func (s *server) Listen() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		client := &Client{
			conn:   conn,
			Server: s,
		}
		go client.listen()
		s.onNewClientCallback(client)
	}
}

// Creates new tcp server instance
func New(address string) *server {
	log.Println("Creating server with address", address)
	server := &server{
		address: address,
	}

	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, response *CommunicationData) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})

	return server
}
