package bambora

type Card struct {
	APIResource
	Token string `json:"token"`
	Code int `json:"code"`
	Version int `json:"version"`
	Message string `json:"message"`
}

type CardParams struct {
	Number *string `form:"number"`
	ExpiryMonth *string `form:"expiry_month"`
	ExpiryYear *string `form:"expiry_year"`
	CVD *string `form:"cvd"`
}