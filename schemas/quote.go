package schemas

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
