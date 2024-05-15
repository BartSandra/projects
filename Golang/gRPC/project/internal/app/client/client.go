package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "project/api/service"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.StreamServiceClient
}

type NumberWithTimestamp struct {
	Number    int32
	Timestamp time.Time
}

func NewClient(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("did not connect to server at %s: %v", address, err)
	}

	client := pb.NewStreamServiceClient(conn)

	return &Client{
		conn:   conn,
		client: client,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Login(username, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.client.Login(ctx, &pb.LoginRequest{Username: username, Password: password})
	if err != nil {
		return fmt.Errorf("could not login: %v", err)
	}

	fmt.Println("Logged in successfully")
	return nil
}

func (c *Client) Stream(interval, duration int) error {
	streamID := uuid.New().String()
	md := metadata.Pairs("uuid", streamID)
	ctx, cancel := context.WithCancel(metadata.NewOutgoingContext(context.Background(), md))
	defer cancel()

	stream, err := c.client.Stream(ctx, &pb.StreamRequest{Interval: int32(interval)})
	if err != nil {
		return fmt.Errorf("could not stream: %v", err)
	}

	var wg sync.WaitGroup
	buffer := make([]NumberWithTimestamp, 0)
	stopRequested := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		startTime := time.Now()
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Printf("Stream ended with error: %v", err)
				return
			}
			buffer = append(buffer, NumberWithTimestamp{Number: res.GetNumber(), Timestamp: time.Unix(res.GetTimestamp(), 0)})

			if time.Since(startTime) >= time.Duration(duration)*time.Second || len(buffer) >= duration {
				if !stopRequested {
					c.Stop(streamID)
					stopRequested = true
				}
				cancel()
				break
			}
		}
	}()

	wg.Wait()

	for _, num := range buffer {
		fmt.Printf("Number: %d, Timestamp: %s\n", num.Number, num.Timestamp.Format(time.RFC3339))
	}

	return nil
}

func (c *Client) Stop(streamID string) error {
	_, err := c.client.Stop(context.Background(), &pb.StopRequest{Uuid: streamID})
	if err != nil {
		return fmt.Errorf("could not stop stream %s: %v", streamID, err)
	}
	fmt.Println("Stream stopped successfully")
	return nil
}

func Run(address, username, password string, interval, duration int) {
	client, err := NewClient(address)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	if err := client.Login(username, password); err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	if err := client.Stream(interval, duration); err != nil {
		log.Fatalf("Failed to stream: %v", err)
	}
}
