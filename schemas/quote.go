package schemas

import "freight-cote-api/schemas/request"

type Address struct {
	ZipCode string `validate:"required" form:"zipcode" json:"zipcode"`
}

type Recipient struct {
	Address Address `validate:"required" form:"address" json:"address"`
}

type Volume struct {
	Category      string  `validate:"required" form:"category" json:"category"`
	Amount        int64   `validate:"required" form:"amount" json:"amount"`
	UnitaryWeight float64 `validate:"required" form:"unitary_weight" json:"unitary_weight"`
	Price         float64 `validate:"required" form:"price" json:"price"`
	Sku           string  `form:"sku" json:"sku,omitempty"`
	Height        float64 `validate:"required" form:"height" json:"height"`
	Width         float64 `validate:"required" form:"width" json:"width"`
	Length        float64 `validate:"required" form:"length" json:"length"`
}

type QuoteInputDTO struct {
	Recipient Recipient `validate:"required" form:"recipient" json:"recipient"`
	Volumes   []Volume  `validate:"required" form:"volumes" json:"volumes"`
}

func (qdto *QuoteInputDTO) SeriealizeInput() request.Request {
	req := new(request.Request)

	shiper := request.Shipper{
		RegisteredNumber: "25438296000158",
		Token:            "1d52a9b6b78cf07b08586152459a5c90",
		PlatformCode:     "5AKVkHqCn",
	}
	req.Shipper = shiper

	recipient := request.Recipient{
		Type:    1,
		Country: "BRA",
		Zipcode: 1311000,
	}
	req.Recipient = recipient

	var volumesToRequest []request.Volume

	for _, volume := range qdto.Volumes {
		volumeAux := request.Volume{
			Category:      volume.Category,
			Amount:        volume.Amount,
			UnitaryWeight: volume.UnitaryWeight,
			UnitaryPrice:  volume.Price,
			SKU:           volume.Sku,
			Height:        volume.Height,
			Width:         volume.Width,
			Length:        volume.Length,
		}

		volumesToRequest = append(volumesToRequest, volumeAux)
	}

	dispatchers := []request.Dispatcher{
		{RegisteredNumber: "25438296000158", Zipcode: 29161376, Volumes: volumesToRequest},
	}
	req.Dispatchers = dispatchers

	req.SimulationType = []int64{0}

	return *req
}
