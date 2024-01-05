package response

type CarrierDTO struct {
	Name string `json:"name"`
}

type DeliveryTime struct {
	Days int64 `json:"days"`
}

type Offer struct {
	Offer        int64        `json:"offer"`
	Carrier      CarrierDTO   `json:"carrier"`
	Service      string       `json:"service"`
	DeliveryTime DeliveryTime `json:"delivery_time"`
	CostPrice    float64      `json:"cost_price"`
	FinalPrice   float64      `json:"final_price"`
}

type Dispatcher struct {
	Offers []Offer `json:"offers"`
}

type FreteRapidoResponseDTO struct {
	Dispatchers []Dispatcher `json:"dispatchers"`
}

func (fr *FreteRapidoResponseDTO) SeriealizeQuoteResponse() *QuoteResponse {
	var carrier []Carrier

	for _, dispatcher := range fr.Dispatchers {
		for _, offer := range dispatcher.Offers {
			carrierAux := Carrier{
				Name:     offer.Carrier.Name,
				Service:  offer.Service,
				Deadline: offer.DeliveryTime.Days,
				Price:    offer.FinalPrice,
			}

			carrier = append(carrier, carrierAux)
		}
	}
	return &QuoteResponse{
		Carrier: carrier,
	}
}
