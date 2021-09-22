package bambora

type Card struct {
	APIResource
	Token string `json:"token"`
	Code int `json:"code"`
	Version int `json:"version"`
	Message string `json:"message"`
}

type CardParams struct {
	Number *string
	ExpiryMonth *string
	ExpiryYear *string
	CVD *string
}