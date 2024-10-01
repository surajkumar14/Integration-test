package integrationtest

import (
	"context"
	"database/sql"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"serviceA/connector"
	"serviceA/models/protomodel"

	grpcroutes "serviceA/router/grpc_routes"

	mocks "serviceA/mocks"

	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var (
	lis             *bufconn.Listener
	grpcServer      *grpc.Server
	SQLMock         sqlmock.Sqlmock
	SQLMockDbClient *sql.DB
	SqlDBClient     *sql.DB
	AerospikeMock   mocks.AerospikeClient
)

func init() {
	//Setup the Mock DB connectors
	InitMockDBConnectors()

	lis = bufconn.Listen(bufSize)
	grpcServer = grpc.NewServer()

	// Register both HTTP and gRPC services for Service A
	protomodel.RegisterGetRatesServiceWithHttpServer(grpcServer, &grpcroutes.GetRatesServiceWithHttpServer{})
	protomodel.RegisterGetRatesServiceWithGrpcServer(grpcServer, &grpcroutes.GetRatesServiceWithGrpcServer{})

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("An error occurred: %v", err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

// Helper function to send HTTP requests and return the result
func performRequest(r http.Handler, method, path string) (*http.Response, []byte, error) {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return resp, body, err
}

func InitMockDBConnectors() {
	InitMockSqlDb()

	// InitSqlDbTOTestCrud()
	// InitMockAerospikeDb()
}

func InitMockSqlDb() {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
		return
	}
	SQLMock = mock
	SQLMockDbClient = db
	connector.SetSqlDBClient(db)
}

// func InitMockAerospikeDb() {
// 	// Create a new mock Aerospike client

// 	AerospikeMock := mocks.NewMockAerospikeClient()
// 	connector.SetAerospikeDBClient(AerospikeMock)
// }
