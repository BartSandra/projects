package main

import (
	"flag"
	"fmt"
	"os"
	"project/internal/app/client"
)

const (
	defaultAddress  = "localhost:50051"
	defaultInterval = 1000
	defaultDuration = 10
)

func main() {
	address := flag.String("address", defaultAddress, "The server address")
	username := flag.String("username", "", "Username for login")
	password := flag.String("password", "", "Password for login")
	interval := flag.Int("interval", defaultInterval, "Interval for stream request in ms")
	duration := flag.Int("duration", defaultDuration, "Duration for the client to receive the stream in seconds")
	flag.Parse()

	if *username == "" || *password == "" {
		fmt.Println("Username and password must be provided")
		os.Exit(1)
	}

	client.Run(*address, *username, *password, *interval, *duration)
}
