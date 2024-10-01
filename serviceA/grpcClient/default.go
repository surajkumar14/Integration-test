package grpcclient

import (
	"errors"
	servcieB "serviceB/models/protomodel"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	GetRatesFromServiceBClient servcieB.GetRatesServiceClient
}

var GRPC_Client *GrpcClients

func InitGrpcServiceClients() {
	InitServiceBGrpcClients()
}

func InitServiceBGrpcClients() {
	serviceBConn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	SetServiceBRatesGrpcClient(servcieB.NewGetRatesServiceClient(serviceBConn))
}

func SetServiceBRatesGrpcClient(client servcieB.GetRatesServiceClient) {
	if GRPC_Client == nil {
		GRPC_Client = &GrpcClients{}
	}
	GRPC_Client.GetRatesFromServiceBClient = client
}

func GetServiceBRatesClient() (servcieB.GetRatesServiceClient, error) {
	if GRPC_Client == nil || GRPC_Client.GetRatesFromServiceBClient == nil {
		return nil, errors.New("serviceB getrates client not initialized")
	}
	return GRPC_Client.GetRatesFromServiceBClient, nil
}
