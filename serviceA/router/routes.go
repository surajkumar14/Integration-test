package router

import (
	"fmt"
	"net/http"
	"serviceA/connector"
	"serviceA/controller"
	grpcclient "serviceA/grpcClient"

	grpcroutes "serviceA/router/grpc_routes"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Setup the Gin router
func SetupHttpServer() *gin.Engine {
	router := gin.Default()

	router.Handle("GET", "service_a/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"value": "pong recied from service A"})
	})

	router.Handle("GET", "service_a/getrates_http", func(c *gin.Context) {
		//this will call service B internally to get the rates using HTTP
		response, err := controller.GetRatesFromServiceBUsingHttp()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	router.Handle("GET", "service_a/getrates_grpc", func(c *gin.Context) {
		//this will call service B internally to get the rates using gRPC
		client, err := grpcclient.GetServiceBRatesClient()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response, err := controller.GetRatesFromServiceBUsingGrpc(c, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	router.Handle("GET", "service_a/getrates_fromdb", func(c *gin.Context) {
		db, err := connector.GetSqlDBClient()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		response, err := controller.GetRatesFromDB(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	router.Handle("GET", "service_a/getrates_fromaerospike", func(c *gin.Context) {
		db, err := connector.GetAerospikeDBClient()
		fmt.Print("get data from aerospike", db, err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		response, err := controller.GetRatesFromAerospike(db)
		fmt.Print("get data from aerospike")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	return router
}

// Setup the GRPC server
func SetupGRPCServer() *grpc.Server {
	gRPCServer := grpc.NewServer()
	grpcroutes.RegisterGRPCServers(gRPCServer)
	return gRPCServer
}
