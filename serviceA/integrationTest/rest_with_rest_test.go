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
	"gopkg.in/h2non/gock.v1"
)

func TestRatesHttpAPIWithHTTP(t *testing.T) {
	defer gock.Off()

	gock.New("http://localhost:5000/service_b/getrates").
		Get("/").
		Reply(200).
		JSON(map[string]string{"rates": "1000"})

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	r := router.SetupHttpServer()

	resp, body, err := performRequest(r, "GET", "/service_a/getrates_http")
	require.NoError(t, err)

	// Parse response body
	var ratesResponse controller.RatesResponse
	require.NoError(t, json.Unmarshal(body, &ratesResponse))

	assert.Equal(t, "1000", ratesResponse.Rates)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Ensure that the mock was called
	assert.True(t, gock.IsDone(), "Not all mocks were called")
}
