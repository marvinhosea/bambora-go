package client

import (
	"github.com/marvinhosea/bambora-go"
	"github.com/marvinhosea/bambora-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApiInit(t *testing.T) {
	merchantId := "383610147"
	profilePasscode := "CE4801A6-EB07-4E0D-8589-4CF612A1"
	paymentPasscode := "e58e7305b052490dA8EB693b4d9aF209"

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

		assert.Nil(t, err)
		assert.NotNil(t, card)
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
