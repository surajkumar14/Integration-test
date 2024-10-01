package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	serviceB "serviceB/models/protomodel"

	"github.com/aerospike/aerospike-client-go"
	_ "github.com/go-sql-driver/mysql"
)

type RatesResponse struct {
	Rates string `json:"rates"`
}

type sqlRow struct {
	Rate string `json:"rate"`
}

func GetRatesFromServiceBUsingHttp() (*RatesResponse, error) {
	resp, err := http.Get("http://localhost:5000/service_b/getrates")
	if err != nil {
		return nil, fmt.Errorf("failed to get rates from service B: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var ratesResponse RatesResponse
	err = json.Unmarshal(body, &ratesResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &ratesResponse, nil
}

func GetRatesFromServiceBUsingGrpc(ctx context.Context, client serviceB.GetRatesServiceClient) (*RatesResponse, error) {
	// Check if the client is initialized
	if client == nil {
		fmt.Print("ServiceBClient is not initialized")
		return nil, fmt.Errorf("ServiceBClient is not initialized")
	}
	request := &serviceB.RatesRequest{}
	response, err := client.GetRates(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("could not get rates: %v", err)
	}

	return &RatesResponse{
		Rates: response.Rates,
	}, nil
}

func GetRatesFromDB(database *sql.DB) (*RatesResponse, error) {

	// Execute the SQL query
	query := "SELECT rate FROM hotel_details WHERE hotel_code = ?"
	row := database.QueryRow(query, "1000")
	var rowResult sqlRow
	err := row.Scan(&rowResult.Rate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no hotel found with code 1000")
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &RatesResponse{
		Rates: rowResult.Rate,
	}, nil
}

func GetRatesFromAerospike(client *aerospike.Client) (*RatesResponse, error) {
	key, err := aerospike.NewKey("test", "hotelrates", "testhotel")
	if err != nil {
		return nil, err
	}
	fmt.Print("get data from aerospike client", client)

	record, err := client.Get(nil, key)
	if err != nil {
		return nil, err
	}

	rate, ok := record.Bins["rate"].(string)
	if !ok {
		return nil, errors.New("rate not found or not a string")
	}
	return &RatesResponse{
		Rates: rate,
	}, nil
}
