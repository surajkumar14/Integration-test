package integrationtest

// func TestRatesHttpAPIWithAerospikeDB(t *testing.T) {
// 	// Set up expected behavior for the mock
// 	key, _ := aerospike.NewKey("test", "hotelrates", "testhotel")
// 	record := &aerospike.Record{
// 		Bins: aerospike.BinMap{"rate": "1000"},
// 	}
// 	AerospikeMock.On("Get", nil, key).Return(record, nil)

// 	// Set Gin to test mode
// 	gin.SetMode(gin.TestMode)
// 	r := router.SetupHttpServer()
// 	resp, body, err := performRequest(r, "GET", "/service_a/getrates_fromaerospike")
// 	require.NoError(t, err)

// 	// Parse response body
// 	// var ratesResponse controller.RatesResponse
// 	fmt.Print("body", string(body))
// 	// require.NoError(t, json.Unmarshal(body, &ratesResponse))
// 	// assert.Equal(t, "1000", ratesResponse.Rates)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	// Ensure all expectations were met
// 	AerospikeMock.AssertExpectations(t)

// }
