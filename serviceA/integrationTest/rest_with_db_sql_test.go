package integrationtest

import (
	"encoding/json"
	"net/http"
	"serviceA/controller"
	"serviceA/router"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRatesHttpAPIWithSqlDB(t *testing.T) {

	// // Define the expected query and result
	query := "SELECT rate FROM hotel_details WHERE hotel_code = \\?"
	rows := SQLMock.NewRows([]string{"rate"}).AddRow("1000")

	// Set up the mock expectations
	SQLMock.ExpectQuery(query).WithArgs("1000").WillReturnRows(rows)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	r := router.SetupHttpServer()
	resp, body, err := performRequest(r, "GET", "/service_a/getrates_fromdb")
	require.NoError(t, err)

	// Parse response body
	var ratesResponse controller.RatesResponse
	require.NoError(t, json.Unmarshal(body, &ratesResponse))

	assert.Equal(t, "1000", ratesResponse.Rates)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Ensure all expectations were met
	err = SQLMock.ExpectationsWereMet()
	require.NoError(t, err)
}
