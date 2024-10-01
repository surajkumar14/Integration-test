package router

import (
	"net/http"
	grpcroutes "serviceB/router/grpc_routes"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Setup the Gin router for HTTP
func SetupHttpServer() *gin.Engine {
	router := gin.Default()

	router.Handle("GET", "/service_b/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"value": "pong recied from service B"})
	})

	router.Handle("GET", "/service_b/getrates", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"rates": "2000_from_serviceB_http"})
	})

	return router
}

// Setup the GRPC server
func SetupGRPCServer() *grpc.Server {
	gRPCServer := grpc.NewServer()
	grpcroutes.RegisterGRPCServers(gRPCServer)
	return gRPCServer
}
