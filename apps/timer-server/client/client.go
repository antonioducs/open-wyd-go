package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/pessoal/wyd/packages/protocol/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start(address string) {
	// Loop to retry connection if it fails initially
	for {
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
		if err == nil {
			handleConnection(conn)
			break
		}
		log.Printf("Failed to connect to %s: %v. Retrying in 5 seconds...", address, err)
		time.Sleep(5 * time.Second)
	}
}

func handleConnection(conn *grpc.ClientConn) {
	defer conn.Close()
	c := pb.NewWydServiceClient(conn)

	fmt.Printf("Timer Server connected to Connect Server\n")

	// Periodically ping
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := c.Ping(ctx, &pb.PingRequest{Message: "Timer Tick"})
		if err != nil {
			log.Printf("could not ping: %v", err)
		} else {
			log.Printf("Response: %s", r.Message)
		}
		cancel()
	}
}
