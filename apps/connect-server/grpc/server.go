package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/pessoal/wyd/packages/protocol/gen"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedWydServiceServer
}

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "Pong: " + req.Message}, nil
}

func (s *server) ValidateToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	// TODO: Validate token with persistence service or cache
	return &pb.TokenResponse{Valid: true, UserId: "user-placeholder"}, nil
}

func Start(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWydServiceServer(s, &server{})
	fmt.Printf("gRPC Server running on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
