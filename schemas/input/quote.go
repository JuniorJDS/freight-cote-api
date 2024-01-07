package input

import (
	"freight-cote-api/configs"
	"freight-cote-api/schemas/request"
	"strconv"
)

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

type Quote struct {
	Recipient Recipient `validate:"required" form:"recipient" json:"recipient"`
	Volumes   []Volume  `validate:"required" form:"volumes" json:"volumes"`
}

func (qdto *Quote) SeriealizeInput() request.Request {
	req := new(request.Request)

	registeredNumber := configs.GetSettings()["REGISTEREDNUMBER"]
	token := configs.GetSettings()["TOKEN"]
	platformCode := configs.GetSettings()["PLATFORMCODE"]

	dispatchersZipCodeString := configs.GetSettings()["DISPATCHERSZIPCODE"]
	dispatchersZipCode, _ := strconv.ParseInt(dispatchersZipCodeString, 10, 64)

	recipientAddressZipCodeString := qdto.Recipient.Address.ZipCode
	recipientAddressZipCode, _ := strconv.ParseInt(recipientAddressZipCodeString, 10, 64)

	shiper := request.Shipper{
		RegisteredNumber: registeredNumber,
		Token:            token,
		PlatformCode:     platformCode,
	}
	req.Shipper = shiper

	recipient := request.Recipient{
		Type:    1,
		Country: "BRA",
		Zipcode: recipientAddressZipCode,
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
		{RegisteredNumber: registeredNumber, Zipcode: dispatchersZipCode, Volumes: volumesToRequest},
	}
	req.Dispatchers = dispatchers

	req.SimulationType = []int64{0}

	return *req
}
