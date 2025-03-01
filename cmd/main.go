package main

import (
	"io"
	"log"
	"net"
	"os"
)

func handleConnection(clientConn net.Conn, dbAddress string) {
	dbConn, err := net.Dial("tcp", dbAddress)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		clientConn.Close()
		return
	}
	defer dbConn.Close()

	// Pipe data between client and RDS
	go io.Copy(dbConn, clientConn)
	io.Copy(clientConn, dbConn)
}

func main() {
	publicIP := os.Getenv("PUBLIC_IP") // e.g., 0.0.0.0:5432
	dbHost := os.Getenv("DB_HOST")     // e.g., your-rds-endpoint:5432

	listener, err := net.Listen("tcp", publicIP)
	if err != nil {
		log.Fatalf("Failed to start proxy: %v", err)
	}
	defer listener.Close()

	log.Printf("Listening on %s and forwarding to %s", publicIP, dbHost)

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(clientConn, dbHost)
	}
}
