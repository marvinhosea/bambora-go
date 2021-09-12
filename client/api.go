package client

import (
	"github.com/marvinhosea/bambora-go/card"
	"github.com/marvinhosea/bambora-go/oauth"
	"github.com/marvinhosea/bambora-go/payment"
)

type Api struct {
	Card *card.Client
	Payment *payment.Client
	Oauth *oauth.Client
}

func (a Api) Init()  {
}

