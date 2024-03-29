package services

import (
	"encoding/json"
	"freight-cote-api/configs"
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
	baseURL := configs.GetSettings()["FRETERAPIDO_API_URL"]
	return &QuoteService{
		requestsServices: *NewRequestsService(
			baseURL,
		),
		quoteRepository: *repositories.NewQuoteRepository(),
	}
}

func (q *QuoteService) Create(quoteInputDTO input.Quote) (*r.QuoteResponse, error) {
	var response r.FreteRapidoResponseDTO

	quoteToRequest := quoteInputDTO.SeriealizeInput()

	quoteInputBytes, err := json.Marshal(quoteToRequest)
	if err != nil {
		log.Printf("Error encoding JSON: %s\n", err.Error())
		return nil, err
	}

	result, err := q.requestsServices.SendRequest(http.MethodPost, quoteInputBytes)
	if err != nil {
		log.Printf("Failed to request external API: %s\n", err.Error())
		return nil, err
	}

	err = json.Unmarshal(result, &response)
	if err != nil {
		log.Printf("Error to unmarshal into a response format: %s\n", err.Error())
		return nil, err
	}

	seriealizedQuote := response.SeriealizeQuoteResponse()
	if len(seriealizedQuote.Carrier) == 0 {
		return seriealizedQuote, nil
	}

	err = q.quoteRepository.Insert(*seriealizedQuote)
	if err != nil {
		return nil, err
	}

	return seriealizedQuote, nil
}

func (q *QuoteService) GetMetrics(lastQuote int64) (*r.Metrics, error) {
	metrics, err := q.quoteRepository.GetMetrics(lastQuote)
	if err != nil {
		return nil, err
	}
	return metrics, nil
}
