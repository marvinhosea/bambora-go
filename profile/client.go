package profile

import (
	"github.com/marvinhosea/bambora-go"
	"github.com/marvinhosea/bambora-go/config"
	"log"
	"net/http"
)

type Client struct {
	E bambora.Endpoint
	Passcode string
}

func New(params *bambora.ProfileParams) (*bambora.Profile, error) {
	return getClient().New(params)
}

func (c Client) New(params *bambora.ProfileParams) (*bambora.Profile, error) {
	p := &bambora.Profile{}
	m := map[string]string{"name": *params.CardName, "code": *params.CardToken}
	err := c.E.Call(http.MethodPost, "/profiles", c.Passcode, map[string]interface{}{
		"token": m,
	}, p)

	return p, err
}

func getClient() *Client {
	log.Println(bambora.EncodedProfilePasscode)
	return &Client{E: bambora.GetEndpoint(config.APIEndpoint), Passcode: bambora.EncodedProfilePasscode}
}
