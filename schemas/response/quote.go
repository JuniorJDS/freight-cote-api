package response

type Carrier struct {
	Name     string  `json:"name" bson:"name"`
	Service  string  `json:"service" bson:"service"`
	Deadline int64   `json:"deadline" bson:"deadline"`
	Price    float64 `json:"price" bson:"price"`
}

type QuoteResponse struct {
	Carrier []Carrier `json:"carrier"`
}
