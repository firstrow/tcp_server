package tcp_server

import (
	"bufio"
	"net"
)

// Client holds info about connection
type Client struct {
	conn   net.Conn
	Server *server
}

// TCP server
type server struct {
	clients                  []*Client
	address                  string // Address to open connection: localhost:9999
	proto                    string
	buf_size                 int
	onNewClientCallback      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, message []byte)
}

// Read client data from channel
func (c *Client) listen() {
	reader := bufio.NewReader(c.conn)
	buf := make([]byte, c.Server.buf_size)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}
		c.Server.onNewMessage(c, buf[:n])
	}
}

// Send message to client
func (c *Client) Send(message []byte) error {
	writer := bufio.NewWriter(c.conn)
	_, err := writer.Write(message)
	if err != nil {
		return err
	}
	return writer.Flush()

}

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
func (s *server) OnNewMessage(callback func(c *Client, message []byte)) {
	s.onNewMessage = callback
}

// Set read buffer size(default 1024 byte)
func (s *server) SetReadBufferSize(size int) {
	s.buf_size = size
}

// Start network server
func (s *server) Listen() error {
	listener, err := net.Listen(s.proto, s.address)
	if err != nil {
		return err
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
func New(proto, address string) *server {
	server := &server{
		address:  address,
		proto:    proto,
		buf_size: 1024,
	}

	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, message []byte) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})

	return server
}
