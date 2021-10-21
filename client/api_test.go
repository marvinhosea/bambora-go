package client

import (
	"github.com/marvinhosea/bambora-go"
	"github.com/marvinhosea/bambora-go/config"
	"github.com/marvinhosea/bambora-go/util"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestApiInit(t *testing.T) {
	merchantId := ""
	profilePasscode := ""
	paymentPasscode := ""

	bc := &Api{}
	bc.Init(merchantId, profilePasscode, paymentPasscode)

	card := &bambora.CardParams{
		Number: util.String("4030000010001234"),
		ExpiryMonth: util.String("02"),
		ExpiryYear: util.String("20"),
		CVD: util.String("123"),
	}

	profile := &bambora.ProfileParams{CardName: util.String("John Doe")}

	t.Run("get card token", func(t *testing.T) {
		card, err := bc.Card.Tokenize(card)
		if err != nil {
			log.Fatalln(err)
		}
		profile.CardToken = &card.Token
		prf, err := bc.Profile.New(profile)
		if err != nil {
			log.Fatalln(err)
		}
		payment, err := bc.Payment.TakePayment(&bambora.PaymentParams{
			Amount: 100,
			PaymentMethod: config.ProfilePaymentMethod,
			Profile: bambora.PaymentProfile{
				CustomerCode: prf.CustomerCode,
				CardId: "1",
				Complete: true,
			},
		})
		assert.Nil(t, err)
		assert.NotNil(t, payment)
	})

	t.Run("get card profile", func(t *testing.T) {
		card, err := bc.Card.Tokenize(card)
		if err != nil {
			t.Errorf("could not create card token %e", err)
		}
		profile.CardToken = util.String(card.Token)

		prfl, err := bc.Profile.New(profile)
		if err != nil {
			t.Errorf("could not create payment profile %e", err)
		}

		assert.NotNil(t, prfl)
	})
}
