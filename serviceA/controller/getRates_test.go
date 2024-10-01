package controller_test

import (
	"context"
	"serviceA/controller"
	"testing"
	"time"

	serviceB "serviceB/models/protomodel"

	mocks "serviceA/mocks"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/h2non/gock.v1"
)

func TestGetRatesHttpService(t *testing.T) {
	// Step 1: Set up `gock` to intercept the request to Service B and return a mock response
	defer gock.Off() // Ensures gock is cleaned up after the test

	gock.New("http://localhost:5000/service_b/getrates").
		Get("/").
		Reply(200).
		JSON(map[string]string{"rates": "1000"})

	// Step 2: Call Service A and check if it properly handles the mocked response from Service B
	response, err := controller.GetRatesFromServiceBUsingHttp()

	assert.NoError(t, err)
	assert.Equal(t, "1000", response.Rates)

	// Step 3: Ensure that the request was indeed intercepted and mocked
	assert.True(t, gock.IsDone(), "Expected all mock requests to be made")
}

func TestGetRatesGrpcService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockGetRatesServiceClient(ctrl)

	// Set up the expected response
	expectedResponse := &serviceB.RatesResponse{
		Rates: "10000",
	}

	mockClient.EXPECT().
		GetRates(gomock.Any(), gomock.Any()).
		Return(expectedResponse, nil)

		// Call the function under test
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := mockClient
	request := &serviceB.RatesRequest{}
	response, err := client.GetRates(ctx, request)
	if err != nil {
		t.Fatalf("could not get rates: %v", err)
	}

	// Assert the response
	assert.Equal(t, expectedResponse.Rates, response.Rates)
}

func TestGetRatesFromDB(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Define the expected query and result
	query := "SELECT rate FROM hotel_details WHERE hotel_code = \\?"
	rows := sqlmock.NewRows([]string{"rate"}).AddRow("1000")

	// Set up the mock expectations
	mock.ExpectQuery(query).WithArgs("1000").WillReturnRows(rows)

	// Call the function
	result, err := controller.GetRatesFromDB(db)
	require.NoError(t, err)

	// Assert the result
	assert.Equal(t, "1000", result.Rates)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
