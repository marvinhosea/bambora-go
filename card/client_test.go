package card

import (
	"github.com/marvinhosea/bambora-go"
	"testing"

	"github.com/marvinhosea/bambora-go/util"
	"github.com/stretchr/testify/assert"
)


func TestCardNew(t *testing.T) {
	merchantId := "383610147"
	profilePasscode := "CE4801A6-EB07-4E0D-8589-4CF612A1"
	paymentPasscode := "e58e7305b052490dA8EB693b4d9aF209"

	_ = bambora.GeneratePasscodes(
		merchantId,
		profilePasscode,
		paymentPasscode)

	token, err := Tokenize(&bambora.CardParams{
		Number: util.String("4030000010001234"),
		ExpiryMonth: util.String("02"),
		ExpiryYear: util.String("20"),
		CVD: util.String("123"),
	})

	assert.Nil(t, err)
	assert.NotNil(t, token)
}