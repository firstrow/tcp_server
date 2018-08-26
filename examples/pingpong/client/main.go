package main

import (
	"encoding/gob"
	"log"
	"net"
	"time"
)

// CommunicationData blah blah blah
type CommunicationData struct {
	Type string
	Data map[string]string
}

func main() {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		log.Println(err)
		Restart()
	}
	for {
		decoder := gob.NewDecoder(conn)
		r := &CommunicationData{}
		err := decoder.Decode(r)
		if err != nil {
			log.Println(err)
			Restart()
		}
		if r.Type == "ping" {
			encoder := gob.NewEncoder(conn)
			r := &CommunicationData{Type: "pong", Data: map[string]string{}}
			encoder.Encode(r)
		}
	}
}

// Restart the client
func Restart() {
	log.Println("Couldn't connect. Retrying...")
	time.Sleep(5 * time.Second)
	main()
}
