package request

type Request struct {
	Shipper        Shipper      `json:"shipper"`
	Recipient      Recipient    `json:"recipient"`
	Dispatchers    []Dispatcher `json:"dispatchers"`
	SimulationType []int64      `json:"simulation_type"`
}

// TODO: create function to manipulate the data to save in mongoDB

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Recipient struct {
	Type    int64  `json:"type"`
	Country string `json:"country"`
	Zipcode int64  `json:"zipcode"`
}

type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          int64    `json:"zipcode"`
	Volumes          []Volume `json:"volumes"`
}

type Volume struct {
	Category      string  `json:"category"`
	Amount        int64   `json:"amount"`
	UnitaryWeight float64 `json:"unitary_weight"`
	UnitaryPrice  float64 `json:"unitary_price"`
	SKU           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}
