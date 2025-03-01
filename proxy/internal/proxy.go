package internal

import (
	"io"
	"log"
	"net"
)

func StartProxy(publicIP, dbHost string) error {
	listener, err := net.Listen("tcp", publicIP)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("Proxy running on %s -> %s", publicIP, dbHost)

	for {
		client, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(client, dbHost)
	}
}

func handleConnection(client net.Conn, dbHost string) {
	rds, err := net.Dial("tcp", dbHost)
	if err != nil {
		log.Printf("Failed to connect to RDS: %v", err)
		client.Close()
		return
	}

	go io.Copy(rds, client)
	io.Copy(client, rds)
	client.Close()
	rds.Close()
}
