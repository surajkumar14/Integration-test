package main

import (
	"log"
	"net"
	"serviceA/connector"
	grpc_client "serviceA/grpcClient"
	"serviceA/router"
)

func init() {
	// Initialize the database
	connector.InitDBConnectors()

	// Initialize the gRPC clients
	grpc_client.InitGrpcServiceClients()

}
func main() {

	// HTTP server (Gin) setup
	httpServer := router.SetupHttpServer()

	// gRPC server setup
	gRPCServer := router.SetupGRPCServer()

	// Start the HTTP server in a goroutine
	go func() {
		if err := httpServer.Run(":3000"); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
		log.Println("Http server is running on port 3000")
	}()

	// Start the gRPC server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", ":3001") // gRPC server on port 5001
		if err != nil {
			log.Fatalf("Failed to listen on port 3001: %v", err)
		}

		log.Println("gRPC server is running on port 3001")
		if err := gRPCServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	select {}
}
