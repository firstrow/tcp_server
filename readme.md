Package TcpServer created to help build simple tcp servers faster.

Here's an example:

```
package sample

import (
	"github.com/firstrow/tcp_server"
)

func main() {
	server := tcp_server.New("localhost:9999")

	server.OnNewClient(func(c *tcp_server.Client) {
		// new client connected
	})
	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		// new messaged from client
	})
	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		// connection with client lost
	})

	go server.Listen()
}
```