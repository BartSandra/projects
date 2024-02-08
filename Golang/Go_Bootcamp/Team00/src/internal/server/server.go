package server

import (
	"github.com/google/uuid"
	"log"
	"math/rand"
	"team00/internal/transmitter/proto"
	"time"
)

type Server struct {
	proto.UnimplementedTransmitterServiceServer
}

func (s *Server) StreamData(req *proto.StreamRequest, stream proto.TransmitterService_StreamDataServer) error {
	sessionID := uuid.New().String()
	for {
		mean := rand.Float64()*20 - 10
		stdDev := rand.Float64()*1.2 + 0.3

		frequency := rand.NormFloat64()*stdDev + mean

		log.Printf("Session ID: %s, Mean: %f, StdDev: %f, Frequency: %f", sessionID, mean, stdDev, frequency)

		if err := stream.Send(&proto.TransmitterData{
			SessionId: sessionID,
			Frequency: frequency,
			Timestamp: time.Now().Unix(),
		}); err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}
}
