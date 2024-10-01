package connector

import (
	"errors"

	"github.com/aerospike/aerospike-client-go"
)

func InitAerospikeDB() {
	// Set up the Aerospike connection
	client, err := aerospike.NewClient("127.0.0.1", 8000)
	if err != nil {
		return
	}
	SetAerospikeDBClient(client)
}

func SetAerospikeDBClient(db *aerospike.Client) {
	if DBClient == nil {
		DBClient = &DataBaseClient{AerospikeClient: db}
	} else {
		DBClient.AerospikeClient = db
	}
}

func GetAerospikeDBClient() (*aerospike.Client, error) {
	if DBClient == nil || DBClient.AerospikeClient == nil {
		return nil, errors.New("DBClient not initialized")
	}
	return DBClient.AerospikeClient, nil
}

func CreateHotelRateRecord(rateKey string, rateValue string) error {
	client, err := GetAerospikeDBClient()
	if err != nil {
		return err
	}

	key, err := aerospike.NewKey("test", "hotelrates", rateKey)
	if err != nil {
		return err
	}

	bins := aerospike.BinMap{
		"rate": rateValue,
	}

	err = client.Put(nil, key, bins)
	if err != nil {
		return err
	}

	return nil
}

func GetHotelRateRecord(rateKey string) (string, error) {
	client, err := GetAerospikeDBClient()
	if err != nil {
		return "", err
	}

	key, err := aerospike.NewKey("test", "hotelrates", rateKey)
	if err != nil {
		return "", err
	}

	record, err := client.Get(nil, key)
	if err != nil {
		return "", err
	}

	rate, ok := record.Bins["rate"].(string)
	if !ok {
		return "", errors.New("rate not found or not a string")
	}

	return rate, nil
}
