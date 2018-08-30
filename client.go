package tcpserver

import "encoding/gob"

// listen listens for data from client
func (c *Client) listen() {
	decoder := gob.NewDecoder(c.conn)
	for {
		r := &communicationData{}
		err := decoder.Decode(r)
		if err != nil {
			c.conn.Close()
			c.Server.onClientDisconnect(c, err)
			return
		}
		c.Server.onNewMessage(c, r.dataType, r.data)
	}
}

// Send sends data to client
func (c *Client) Send(dataType string, data map[string]string) {
	encoder := gob.NewEncoder(c.conn)
	r := &communicationData{dataType, data}
	err := encoder.Encode(r)
	if err != nil {
		c.conn.Close()
		c.Server.onClientDisconnect(c, err)
		return
	}
}
