package grpcroutes

import (
	"context"

	"github.com/surajkumar14/Integration-test/models/protomodel"

	"google.golang.org/grpc"
)

func RegisterGRPCServers(gRPCServer *grpc.Server) {
	protomodel.RegisterGetRatesServiceServer(gRPCServer, &GetRatesServiceServer{})

}

type GetRatesServiceServer struct {
	protomodel.UnimplementedGetRatesServiceServer
}

func (s *GetRatesServiceServer) GetRates(ctx context.Context, request *protomodel.RatesRequest) (*protomodel.RatesResponse, error) {
	return &protomodel.RatesResponse{
		Rates: "1000_from_serviceB_grpc",
	}, nil

}
