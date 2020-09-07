package main

import (
	"log"
	"net"
)

func main() {

	server := newServer()
	go server.processCommandsOnServer()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		go server.handleClientRequest(&conn)
	}
}
