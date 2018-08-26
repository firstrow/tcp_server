package main

import (
	"log"

	"github.com/89apt89/tcpserver"
)

func main() {
	server := tcpserver.New("localhost:2000")

	server.OnNewClient(func(c *tcpserver.Client) {
		log.Printf("Playing ping pong with %s", c.Conn())
		c.Send("ping")
	})
	server.OnNewMessage(func(c *tcpserver.Client, response *tcpserver.CommunicationData) {
		log.Println(response.Type)
	})
	server.OnClientConnectionClosed(func(c *tcpserver.Client, err error) {
		// connection with client lost
	})

	server.Listen()
}
