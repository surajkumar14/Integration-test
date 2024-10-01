package connector

import (
	"database/sql"

	"github.com/aerospike/aerospike-client-go"
)

type DataBaseClient struct {
	SqlDBClient     *sql.DB
	AerospikeClient *aerospike.Client
}

var DBClient *DataBaseClient

func InitDBConnectors() {
	// Initialize the database
	InitSqlDataBase()

	// Initialize the Aerospike database
	InitAerospikeDB()
}
