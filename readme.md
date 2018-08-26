# TCPServer
A TCP server implementation using gobs for communication

### Install package

``` bash
go get -u github.com/89apt89/tcpserver
```

### Usage:

``` go
package main

import "github.com/89apt89/tcpserver"

func main() {
	server := tcp_server.New("localhost:9999")

	server.OnNewClient(func(c *tcpserver.Client) {
		// new client connected
		// lets send some message
		c.Send("Hello")
	})
	server.OnNewMessage(func(c *tcpserver.Client, response *tcpserver.Data) {
		// new message received
	})
	server.OnClientConnectionClosed(func(c *tcpserver.Client, err error) {
		// connection with client lost
	})

	server.Listen()
}
```

# Contributing

To hack on this project:

1. Install as usual (`go get -u github.com/89apt89/tcpserver`)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Ensure everything works and the tests pass (`go test`)
4. Commit your changes (`git commit -am 'Add some feature'`)

Contribute upstream:

1. Fork it on GitHub
2. Add your remote (`git remote add fork git@github.com:89apt89/tcpserver.git`)
3. Push to the branch (`git push fork my-new-feature`)
4. Create a new Pull Request on GitHub

Notice: Always use the original import path by installing with `go get`.
