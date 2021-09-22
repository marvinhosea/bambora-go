package payment

import (
	"github.com/marvinhosea/bambora-go"
	card2 "github.com/marvinhosea/bambora-go/card"
	"github.com/marvinhosea/bambora-go/config"
	profile2 "github.com/marvinhosea/bambora-go/profile"
	"github.com/marvinhosea/bambora-go/util"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTakePayment(t *testing.T) {
	t.Run("take payment", func(t *testing.T) {
		merchantId := "383610147"
		profilePasscode := "e58e7305b052490dA8EB693b4d9aF209"
		paymentPasscode := "91E6726a7aA5499AA5AFa85C668093CD"

		cnf := bambora.GeneratePasscodes(
			merchantId,
			profilePasscode,
			paymentPasscode)

		bambora.EncodedPaymentPasscode = cnf.EPaymentPasscode
		bambora.EncodedProfilePasscode = cnf.EProfilePasscode

		card, err := card2.Tokenize(&bambora.CardParams{
			Number: util.String("4030000010001234"),
			ExpiryMonth: util.String("02"),
			ExpiryYear: util.String("20"),
			CVD: util.String("123"),
		})

		profile, err := profile2.New(&bambora.ProfileParams{
			CardName: util.String("John Doe"),
			CardToken: util.String(card.Token),
		})

		log.Println(profile, "profile")

		payment, err := TakePayment(&bambora.PaymentParams{
			Amount: 100,
			PaymentMethod: config.ProfilePaymentMethod,
			Profile: bambora.PaymentProfile{
				CustomerCode: profile.CustomerCode,
				CardId: "1",
				Complete: true,
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, payment)
	})
}