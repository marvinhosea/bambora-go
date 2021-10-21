package profile

import (
	"github.com/marvinhosea/bambora-go"
	card2 "github.com/marvinhosea/bambora-go/card"
	"github.com/marvinhosea/bambora-go/util"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestClientNew(t *testing.T) {
	merchantId := ""
	profilePasscode := ""
	paymentPasscode := ""

	con := bambora.GeneratePasscodes(
		merchantId,
		profilePasscode,
		paymentPasscode)
	bambora.EncodedProfilePasscode = con.EProfilePasscode
	bambora.EncodedPaymentPasscode = con.EPaymentPasscode

	card, err := card2.Tokenize(&bambora.CardParams{
		Number: util.String("4030000010001234"),
		ExpiryMonth: util.String("02"),
		ExpiryYear: util.String("20"),
		CVD: util.String("123"),
	})

	profile, err := New(&bambora.ProfileParams{
		CardName: util.String("John Doe"),
		CardToken: util.String(card.Token),
	})

	log.Println(profile.CustomerCode)

	assert.Nil(t, err)
	assert.NotNil(t, profile)
}
