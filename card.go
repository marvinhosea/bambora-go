package bambora

type Card struct {
	APIResource
}


type CardParams struct {
	Number *string `form:"number"`
	ExpiryMonth *string `form:"expiry_month"`
	ExpiryYear *string `form:"expiry_year"`
	CVD *string `form:"cvd"`
}