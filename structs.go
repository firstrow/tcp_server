package tcpserver

import "net"

// Client stores the data of a connected client
type Client struct {
	conn   net.Conn
	Server *server
}

type server struct {
	address            string
	onNewClient        func(c *Client)
	onNewMessage       func(c *Client, dataType string, data map[string]string)
	onClientDisconnect func(c *Client, err error)
}

type communicationData struct {
	dataType string
	data     map[string]string
}
