package connector

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitSqlDataBase() error {
	// Set up the MySQL connection
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("Failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("Failed to open database: %v", err)
	}
	return nil
}

func TestSQLCRUDOperations(t *testing.T) {
	err := InitSqlDataBase()
	assert.NoError(t, err)

}
