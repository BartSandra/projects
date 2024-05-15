package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "project/api/service"
)

type server struct {
	pb.UnimplementedStreamServiceServer
	activeStreams sync.Map
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Printf("Username: %s, Password: %s\n", in.GetUsername(), in.GetPassword())
	return &pb.LoginResponse{Message: "Logged in successfully"}, nil
}

func (s *server) Stream(in *pb.StreamRequest, stream pb.StreamService_StreamServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("could not get metadata from context")
	}

	uuids := md.Get("uuid")
	if len(uuids) == 0 {
		return fmt.Errorf("no uuid in metadata")
	}

	streamID := uuids[0]

	ticker := time.NewTicker(time.Duration(in.GetInterval()) * time.Millisecond)
	defer ticker.Stop()

	s.activeStreams.Store(streamID, true)

	for i := 1; ; i++ {
		select {
		case <-stream.Context().Done():
			s.activeStreams.Delete(streamID)
			log.Println("Stream was cancelled.")
			return nil
		case <-ticker.C:
			isActive, ok := s.activeStreams.Load(streamID)
			if !ok {
				return fmt.Errorf("stream not found")
			}
			if !isActive.(bool) {
				return nil
			}
			if err := stream.Send(&pb.StreamResponse{Number: int32(i), Timestamp: time.Now().Unix()}); err != nil {
				return err
			}
		}
	}
}

func (s *server) Stop(ctx context.Context, in *pb.StopRequest) (*pb.StopResponse, error) {
	streamID := in.Uuid
	if _, ok := s.activeStreams.Load(streamID); ok {
		s.activeStreams.Store(streamID, false)
		fmt.Println("Stop request received for:", streamID)
		return &pb.StopResponse{Message: "Stopped successfully"}, nil
	}
	return nil, fmt.Errorf("stream not found")
}

func Run(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	s := &server{}
	pb.RegisterStreamServiceServer(srv, s)
	fmt.Printf("Server listening at %v\n", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
