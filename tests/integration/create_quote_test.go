package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateQuote__simulateQuoteWithValidBodyAndValidResponseFromFreteAPI__expectedSuccessAndResultSavedInDatabase(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// FIXTURE
	baseURL := "https://test-frete-api/api/v3/quote/simulate"

	httpmock.RegisterResponder("POST", baseURL,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, httpmock.File("data/body.json"))
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, err
		},
	)

	quoteInput := map[string]interface{}{
		"recipient": map[string]interface{}{
			"address": map[string]string{
				"zipcode": "01311000",
			},
		},
		"volumes": []map[string]interface{}{
			{
				"category":       "7",
				"amount":         1,
				"unitary_weight": 5,
				"price":          349,
				"sku":            "abc-teste-123",
				"height":         0.2,
				"width":          0.2,
				"length":         0.2,
			},
			{
				"category":       "7",
				"amount":         2,
				"unitary_weight": 4,
				"price":          556,
				"sku":            "abc-teste-527",
				"height":         0.4,
				"width":          0.6,
				"length":         0.15,
			},
		},
	}

	jsonBytes, err := json.Marshal(quoteInput)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	quoteInputBody := bytes.NewBuffer(jsonBytes)

	// EXERCISE
	url := "/api/v1/quote"
	resp, errResp := baseTest.appClient("POST", url, quoteInputBody)

	var respData map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(respBody, &respData)

	// ASSERTS
	assert.NoError(t, errResp)

	expectedRespData := map[string]interface{}{
		"carrier": []interface{}{
			map[string]interface{}{"name": "JADLOG", "service": ".PACKAGE", "deadline": float64(13), "price": 35.99},
			map[string]interface{}{"name": "CORREIOS", "service": "PAC", "deadline": float64(15), "price": 44.96},
			map[string]interface{}{"name": "CORREIOS", "service": "SEDEX", "deadline": float64(11), "price": 74.17},
			map[string]interface{}{"name": "BTU BRASPRESS", "service": "Normal", "deadline": float64(15), "price": 93.35},
			map[string]interface{}{"name": "CORREIOS", "service": "PAC", "deadline": float64(15), "price": 112.96},
			map[string]interface{}{"name": "CORREIOS", "service": "SEDEX", "deadline": float64(11), "price": 205.54},
			map[string]interface{}{"name": "PRESSA FR (TESTE)", "service": "Normal", "deadline": float64(11), "price": 1599.39},
			map[string]interface{}{"name": "PRESSA FR (TESTE)", "service": "Normal", "deadline": float64(11), "price": 1599.39},
		},
	}

	assert.EqualValues(t, expectedRespData, respData)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	quoteCollection := baseTest.database.Collection("quote")
	cursor, _ := quoteCollection.Find(ctx, bson.M{})

	carriers := []map[string]interface{}{}
	err = cursor.All(ctx, &carriers)
	if err != nil {
		log.Printf("Error to fetch carriers: %s\n", err.Error())
	}

	assert.Len(t, carriers, 8)
}

func TestCreateQuote__simulateQuoteWithValidBodyAndEmptyOffersFromFreteAPI__expectedSuccessAndReturnEmptyCarrier(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// FIXTURE
	baseURL := "https://test-frete-api/api/v3/quote/simulate"

	httpmock.RegisterResponder("POST", baseURL,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, httpmock.File("data/body_with_no_offers.json"))
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, err
		},
	)

	quoteInput := map[string]interface{}{
		"recipient": map[string]interface{}{
			"address": map[string]string{
				"zipcode": "01311000",
			},
		},
		"volumes": []map[string]interface{}{
			{
				"category":       "7",
				"amount":         1,
				"unitary_weight": 5,
				"price":          349,
				"sku":            "abc-teste-123",
				"height":         0.2,
				"width":          0.2,
				"length":         0.2,
			},
			{
				"category":       "7",
				"amount":         2,
				"unitary_weight": 4,
				"price":          556,
				"sku":            "abc-teste-527",
				"height":         0.4,
				"width":          0.6,
				"length":         0.15,
			},
		},
	}

	jsonBytes, err := json.Marshal(quoteInput)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	quoteInputBody := bytes.NewBuffer(jsonBytes)

	// EXERCISE
	url := "/api/v1/quote"
	resp, errResp := baseTest.appClient("POST", url, quoteInputBody)

	var respData map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(respBody, &respData)

	// ASSERTS
	assert.NoError(t, errResp)

	expectedRespData := map[string]interface{}{
		"carrier": []interface{}{},
	}

	assert.EqualValues(t, expectedRespData, respData)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	quoteCollection := baseTest.database.Collection("quote")
	cursor, _ := quoteCollection.Find(ctx, bson.M{})

	carriers := []map[string]interface{}{}
	err = cursor.All(ctx, &carriers)
	if err != nil {
		log.Printf("Error to fetch carriers: %s\n", err.Error())
	}
	assert.Empty(t, carriers)
}
