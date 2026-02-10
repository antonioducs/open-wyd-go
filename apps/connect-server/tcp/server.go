package tcp

import (
	"fmt"
	"log"
	"net"
)

func Start(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to start TCP server: %v", err)
	}
	fmt.Printf("TCP Server running on port %s\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	// TODO: implement packet handling logic
	// e.g., read header, dispatch handler
}
