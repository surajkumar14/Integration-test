package grpcroutes

import (
	"context"
	"fmt"
	"serviceA/controller"
	grpcclient "serviceA/grpcClient"
	"serviceA/models/protomodel"

	"google.golang.org/grpc"
)

func RegisterGRPCServers(gRPCServer *grpc.Server) {
	protomodel.RegisterGetRatesServiceWithGrpcServer(gRPCServer, &GetRatesServiceWithGrpcServer{})
	protomodel.RegisterGetRatesServiceWithHttpServer(gRPCServer, &GetRatesServiceWithHttpServer{})
}

type GetRatesServiceWithGrpcServer struct {
	protomodel.UnimplementedGetRatesServiceWithGrpcServer
}

type GetRatesServiceWithHttpServer struct {
	protomodel.UnimplementedGetRatesServiceWithHttpServer
}

func (s *GetRatesServiceWithGrpcServer) GetRatesGrpc(ctx context.Context, request *protomodel.RatesRequestGrpc) (*protomodel.RatesResponseGrpc, error) {
	client, err := grpcclient.GetServiceBRatesClient()
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}
	response, err := controller.GetRatesFromServiceBUsingGrpc(ctx, client)
	if err != nil {
		return nil, err
	}
	return &protomodel.RatesResponseGrpc{
		Rates: response.Rates,
	}, nil
}

func (s *GetRatesServiceWithHttpServer) GetRatesHttp(ctx context.Context, request *protomodel.RatesRequestHttp) (*protomodel.RatesResponseHttp, error) {
	response, err := controller.GetRatesFromServiceBUsingHttp()
	if err != nil {
		return nil, err
	}
	return &protomodel.RatesResponseHttp{
		Rates: response.Rates,
	}, nil
}
