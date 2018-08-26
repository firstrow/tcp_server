package main

import (
	"encoding/gob"
	"fmt"
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
	decoder := gob.NewDecoder(conn)
	for {
		r := &CommunicationData{}
		err := decoder.Decode(r)
		if err != nil {
			log.Println(err)
			Restart()
		}
		log.Println(r.Type)
		if r.Type == "ping" {
			fmt.Printf("Ping from %s\n", conn.RemoteAddr())
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
