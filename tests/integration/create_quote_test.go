package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateQuote__createQuoteWithValidBody__expectedSuccess(t *testing.T) {
	baseTest := NewBaseTest()

	// FIXTURE
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

}
