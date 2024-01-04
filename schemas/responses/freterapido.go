package responses

import "time"

type Carrier struct {
	Name             string `json:"name"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Logo             string `json:"logo"`
	Reference        int    `json:"reference"`
	CompanyName      string `json:"company_name"`
}

type DeliveryTime struct {
	Days          int    `json:"days"`
	Hours         int    `json:"hours,omitempty"`
	Minutes       int    `json:"minutes,omitempty"`
	EstimatedDate string `json:"estimated_date"`
}

type Weights struct {
	Real  float64 `json:"real"`
	Cubed float64 `json:"cubed,omitempty"`
	Used  float64 `json:"used,omitempty"`
}

type Correios struct {
	DeclaredValue bool `json:"declared_value"`
}

type Offer struct {
	Offer                int64        `json:"offer"`
	TableReference       string       `json:"table_reference,omitempty"`
	SimulationType       int64        `json:"simulation_type"`
	Carrier              Carrier      `json:"carrier"`
	Service              string       `json:"service"`
	ServiceCode          string       `json:"service_code,omitempty"`
	ServiceDescription   string       `json:"service_description,omitempty"`
	DeliveryTime         DeliveryTime `json:"delivery_time"`
	Expiration           time.Time    `json:"expiration"`
	CostPrice            float64      `json:"cost_price"`
	FinalPrice           float64      `json:"final_price"`
	Weights              Weights      `json:"weights"`
	OriginalDeliveryTime DeliveryTime `json:"original_delivery_time"`
	Identifier           string       `json:"identifier,omitempty"`
	Correios             Correios     `json:"correios,omitempty"`
	HomeDelivery         bool         `json:"home_delivery"`
}

type Dispatcher struct {
	ID                         string  `json:"id"`
	RequestID                  string  `json:"request_id"`
	RegisteredNumberShipper    string  `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string  `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int     `json:"zipcode_origin"`
	Offers                     []Offer `json:"offers"`
}

type FreteRapidoResponse struct {
	Dispatchers []Dispatcher `json:"dispatchers"`
}
