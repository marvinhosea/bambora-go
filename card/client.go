package card

import (
	go_bambora "github.com/marvinhosea/bambora-go"
	"github.com/marvinhosea/bambora-go/config"
	"net/http"
)

type Client struct {
	E go_bambora.Endpoint
	Passcode string
}

func New(params *go_bambora.CardParams) (*go_bambora.Card, error) {
	return getClient().Tokenize(params)
}

func (c Client) Tokenize(params *go_bambora.CardParams) (*go_bambora.Card, error) {
	APIResource := go_bambora.APIResource{LastResponse: &go_bambora.ApiResponse{Status: "demo"}}
	p := &go_bambora.Card{
		APIResource,
	}
	err := c.E.Call(http.MethodPost, "/dmeo", c.Passcode, nil, p)
	return p, err
}

func getClient() *Client {
	return &Client{go_bambora.GetEndpoint(config.APIEndpoint), go_bambora.AccountPasscode}
}