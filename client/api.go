package client

import (
	"github.com/marvinhosea/bambora-go"
	"github.com/marvinhosea/bambora-go/card"
	"github.com/marvinhosea/bambora-go/config"
	"github.com/marvinhosea/bambora-go/payment"
	"github.com/marvinhosea/bambora-go/profile"
	"log"
)

type Api struct {
	Card *card.Client
	Payment *payment.Client
	Profile *profile.Client
}

func (a *Api) Init(merchantId, profilePasscode string, paymentProfile string)  {
	conf := bambora.GeneratePasscodes(merchantId, profilePasscode, paymentProfile)
	if conf == nil {
		log.Fatalln("failed to create config")
	}

	endpoints := &bambora.Endpoints{
		API: bambora.GetEndpoint(config.APIEndpoint),
		Connect: bambora.GetEndpoint(config.Connect),
	}

	a.Card = &card.Client{E: endpoints.Connect, Passcode: conf.EProfilePasscode}
	a.Profile = &profile.Client{E: endpoints.API, Passcode: conf.EProfilePasscode}
	a.Payment = &payment.Client{E: endpoints.API, Passcode: conf.EPaymentPasscode}
}

