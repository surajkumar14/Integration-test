package integrationtest

import (
	"context"
	"serviceA/models/protomodel"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"gopkg.in/h2non/gock.v1"
)

func TestRatesGrpcApiWithHttp(t *testing.T) {
	defer gock.Off()

	gock.New("http://localhost:5000/service_b/getrates").
		Get("/").
		Reply(200).
		JSON(map[string]string{"rates": "1000"})

	// gRPC client setup for Service A
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	client := protomodel.NewGetRatesServiceWithHttpClient(conn)

	// Make a gRPC call to Service A
	resp, err := client.GetRatesHttp(context.Background(), &protomodel.RatesRequestHttp{})
	require.NoError(t, err)

	assert.Equal(t, "1000", resp.Rates)
}
