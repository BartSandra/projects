package main

import (
	"flag"
	"fmt"
	"project/internal/app/server"
)

const (
	defaultPort = 50051
)

func main() {
	port := flag.Int("port", defaultPort, "The server port")
	flag.Parse()

	if *port == 0 {
		fmt.Println("Valid port must be specified")
		flag.Usage()
		return
	}

	server.Run(*port)
}
