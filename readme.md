# TCPServer
Package tcp_server created to help build TCP servers faster.

### Install package

``` bash
> go get github.com/firstrow/tcp_server
```

### Usage:

NOTICE: `OnNewMessage` callback will receive new message only if it's ending with `\n`

``` go
package main

import "github.com/firstrow/tcp_server"
import "log"

func main() {
	server := tcp_server.New("tcp4", "localhost:9999")
	server.SetReadBufferSize(64)
	server.OnNewClient(func(c *tcp_server.Client) {
		// new client connected
		// lets send some message
		c.Send([]byte("Hello"))
	})
	server.OnNewMessage(func(c *tcp_server.Client, message []byte) {
		// new message received
	})
	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		// connection with client lost
	})

	err := server.Listen()
	if err != nil {
		log.Panicln(err)
	}
}
```

# Contributing

To hack on this project:

1. Install as usual (`go get -u github.com/firstrow/tcp_server`)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Ensure everything works and the tests pass (`go test`)
4. Commit your changes (`git commit -am 'Add some feature'`)

Contribute upstream:

1. Fork it on GitHub
2. Add your remote (`git remote add fork git@github.com:firstrow/tcp_server.git`)
3. Push to the branch (`git push fork my-new-feature`)
4. Create a new Pull Request on GitHub

Notice: Always use the original import path by installing with `go get`.
