package response

type Metric struct {
	Name     string  `json:"name" bson:"name"`
	Quantity int64   `json:"quantity" bson:"quantity"`
	Total    float64 `json:"total" bson:"total"`
	AVG      float64 `json:"avg" bson:"avg"`
}

type Metrics struct {
	ByCarriers   []Metric `json:"by_carriers" bson:"by_carriers"`
	HighestPrice float64  `json:"highest_price" bson:"highest_price"`
	LowestPrice  float64  `json:"lowest_price" bson:"lowest_price"`
}
