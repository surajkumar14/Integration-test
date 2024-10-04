package main

import (
	"log"
	"net"

	"github.com/surajkumar14/ServiceB/router"
)

func main() {

	// HTTP server (Gin) setup
	httpServer := router.SetupHttpServer()

	// gRPC server setup
	gRPCServer := router.SetupGRPCServer()

	// Start the HTTP server in a goroutine
	go func() {
		if err := httpServer.Run(":5000"); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
		log.Println("Http server is running on port 5000")
	}()

	// Start the gRPC server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", ":5001") // gRPC server on port 5001
		if err != nil {
			log.Fatalf("Failed to listen on port 5001: %v", err)
		}

		log.Println("gRPC server is running on port 5001")
		if err := gRPCServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	select {}
}
