package integration_test

import (
	"encoding/json"
	"freight-cote-api/schemas/response"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQuoteMetrics__GetMetricsFromDatabaseQuoteCollection__ExpectedSuccess(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	// FIXTURE
	_ = baseTest.populateCollectionDatabase("quote", "./data/quote_collection.json")

	// EXERCISE
	url := "/api/v1/quote/metrics"
	resp, errResp := baseTest.appClient("GET", url, nil)

	var respData response.Metrics
	respBody, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(respBody, &respData)

	// ASSERTS
	assert.NoError(t, errResp)

	expectedData := response.Metrics{
		ByCarriers: []response.Metric{
			{Name: "BTU BRASPRESS", Quantity: 1, Total: 90, AVG: 90},
			{Name: "JADLOG", Quantity: 1, Total: 20, AVG: 20},
			{Name: "CORREIOS", Quantity: 4, Total: 290, AVG: 72.5},
			{Name: "PRESSA FR (TESTE)", Quantity: 2, Total: 4000, AVG: 2000},
		},
		HighestPrice: 3000,
		LowestPrice:  20,
	}
	assert.EqualValues(t, expectedData, respData)

}

func TestGetQuoteMetrics__GetMetricsFromDatabaseQuoteCollectionWithLastQuotesQueryParameter__ExpectedSuccess(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	// FIXTURE
	_ = baseTest.populateCollectionDatabase("quote", "./data/quote_collection.json")

	// EXERCISE
	url := "/api/v1/quote/metrics?last_quotes=2"
	resp, errResp := baseTest.appClient("GET", url, nil)

	var respData response.Metrics
	respBody, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(respBody, &respData)

	// ASSERTS
	assert.NoError(t, errResp)

	expectedData := response.Metrics{
		ByCarriers: []response.Metric{
			{Name: "PRESSA FR (TESTE)", Quantity: 2, Total: 4000, AVG: 2000},
		},
		HighestPrice: 3000,
		LowestPrice:  1000,
	}
	assert.EqualValues(t, expectedData, respData)

}

func TestGetQuoteMetrics__GetMetricsFromDatabaseQuoteCollectionWithoutData__ExpectedSuccessAndEmptyMetrics(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	// EXERCISE
	url := "/api/v1/quote/metrics?last_quotes=8"
	resp, errResp := baseTest.appClient("GET", url, nil)

	var respData response.Metrics
	respBody, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(respBody, &respData)

	// ASSERTS
	assert.NoError(t, errResp)

	expectedData := response.Metrics{
		ByCarriers:   nil,
		HighestPrice: 0,
		LowestPrice:  0,
	}
	assert.EqualValues(t, expectedData, respData)

}
