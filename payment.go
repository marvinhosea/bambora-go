package bambora

type PaymentParams struct {
	Amount float32
	PaymentMethod string
	Profile PaymentProfile
}

type Payment struct {
	APIResource
}

type PaymentProfile struct {
	CustomerCode string `json:"customer_code"`
	CardId string `json:"card_id"`
	Complete bool `json:"complete"`
}