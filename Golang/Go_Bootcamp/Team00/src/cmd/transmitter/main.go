package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"team00/internal/consts"
	"team00/internal/server"
	pb "team00/internal/transmitter/proto"
)

func main() {
	lis, err := net.Listen(consts.Network, consts.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTransmitterServiceServer(s, &server.Server{})
	log.Println("Server started on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
