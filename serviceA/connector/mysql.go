package connector

import (
	"database/sql"
	"errors"
)

func InitSqlDataBase() {
	// Set up the MySQL connection
	dsn := "root:root@123@tcp(127.0.0.1:3306)/local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	SetSqlDBClient(db)
}

func SetSqlDBClient(db *sql.DB) {
	if DBClient == nil {
		DBClient = &DataBaseClient{SqlDBClient: db}
	} else {
		DBClient.SqlDBClient = db
	}
}

func GetSqlDBClient() (*sql.DB, error) {
	if DBClient == nil || DBClient.SqlDBClient == nil {
		return nil, errors.New("DBClient not initialized")
	}
	return DBClient.SqlDBClient, nil
}
