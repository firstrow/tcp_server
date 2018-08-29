package tcpserver

var clients []Client

// OnNewClient called when a new client connects
func (s *server) OnNewClient(callback func(c *Client)) {
	s.onNewClient = callback
}

// OnNewMessage called when the client sends data
func (s *server) OnNewMessage(callback func(c *Client, dataType string, data map[string]string)) {
	s.onNewMessage = callback
}

// OnClientDisconnect called when client disconnects
func (s *server) OnClientDisconnect(callback func(c *Client, err error)) {
	s.onClientDisconnect = callback
}
