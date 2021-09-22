package card

import (
	gobambora "github.com/marvinhosea/bambora-go"
	"github.com/marvinhosea/bambora-go/config"
	"net/http"
)

type Client struct {
	E        gobambora.Endpoint
	Passcode string
}

func Tokenize(params *gobambora.CardParams) (*gobambora.Card, error) {
	return getClient().Tokenize(params)
}

func (c Client) Tokenize(params *gobambora.CardParams) (*gobambora.Card, error) {
	p := &gobambora.Card{}
	err := c.E.Call(http.MethodPost, "/tokenization", c.Passcode, map[string]interface{}{
		"number": *params.Number,
		"expiry_month": *params.ExpiryMonth,
		"expiry_year": *params.ExpiryYear,
		"cvd": *params.CVD,
	}, p)
	return p, err
}

func getClient() *Client {
	return &Client{gobambora.GetEndpoint(config.Connect), gobambora.EncodedProfilePasscode}
}