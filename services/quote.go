package services

import (
	"encoding/json"
	"fmt"
	"freight-cote-api/schemas"
	"freight-cote-api/schemas/responses"
	"log"
	"net/http"
)

type QuoteService struct {
	requestsServices RequestsServices
}

func NewQuoteService() *QuoteService {
	return &QuoteService{
		requestsServices: *NewRequestsService(
			"https://sp.freterapido.com/api/v3/quote/simulate",
		),
	}
}

func (q *QuoteService) QuoteFreightV3(quoteInputDTO schemas.QuoteInputDTO) (*responses.FreteRapidoResponse, error) {
	var response responses.FreteRapidoResponse

	quoteToRequest := quoteInputDTO.SeriealizeInput()

	quoteInputBytes, err := json.Marshal(quoteToRequest)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return nil, err
	}

	result, err := q.requestsServices.SendRequest(http.MethodPost, quoteInputBytes)
	if err != nil {
		log.Printf("Failed to get fs data: %s\n", err.Error())
		return nil, err
	}

	_ = json.Unmarshal(result, &response)

	return &response, nil
}
