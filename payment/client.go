package payment

import (
	"github.com/marvinhosea/bambora-go"
	"github.com/marvinhosea/bambora-go/config"
	"net/http"
)

type Client struct {
	E bambora.Endpoint
	Passcode string
}

func TakePayment(params *bambora.PaymentParams) (*bambora.Payment, error) {
	return getClient().TakePayment(params)
}

func (c *Client) TakePayment(params *bambora.PaymentParams) (*bambora.Payment, error) {
	payment := &bambora.Payment{}
	pma := map[string]interface{}{
		"amount": 100,
		"payment_method": params.PaymentMethod,
		"payment_profile": map[string]interface{}{
			"customer_code": params.Profile.CustomerCode,
			"card_id": params.Profile.CardId,
			"complete": params.Profile.Complete,
		},
	}
	err := c.E.Call(http.MethodPost, "/payments", c.Passcode, pma, payment)

	return payment, err
}

func getClient() *Client {
	return &Client{E: bambora.GetEndpoint(config.APIEndpoint), Passcode: bambora.EncodedPaymentPasscode}
}