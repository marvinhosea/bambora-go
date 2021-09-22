package bambora

type Profile struct {
	APIResource
	CustomerCode string `json:"customer_code"`
}

type ProfileParams struct {
	CardName *string
	CardToken *string
}

