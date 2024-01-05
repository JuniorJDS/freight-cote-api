package services

import (
	"encoding/json"
	"fmt"
	"freight-cote-api/repositories"
	"freight-cote-api/schemas/input"
	r "freight-cote-api/schemas/response"
	"log"
	"net/http"
)

type QuoteService struct {
	requestsServices RequestsServices
	quoteRepository  repositories.QuoteRepository
}

func NewQuoteService() *QuoteService {
	return &QuoteService{
		requestsServices: *NewRequestsService(
			"https://sp.freterapido.com/api/v3/quote/simulate",
		),
		quoteRepository: *repositories.NewQuoteRepository(),
	}
}

func (q *QuoteService) Create(quoteInputDTO input.Quote) (*r.QuoteResponse, error) {
	var response r.FreteRapidoResponseDTO

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
	seriealizedQuote := response.SeriealizeQuoteResponse()

	// TODO: insert response to database
	err = q.quoteRepository.Create(*seriealizedQuote)
	if err != nil {
		return nil, err
	}

	return seriealizedQuote, nil
}
